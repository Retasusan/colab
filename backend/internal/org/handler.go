package org

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	appauth "github.com/Retasusan/colab_backend/internal/auth"
	"gorm.io/gorm"
)

type Organization struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Slug      string    `gorm:"column:slug" json:"slug"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (Organization) TableName() string {
	return "organizations"
}

type Membership struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrganizationID int64     `gorm:"column:organization_id" json:"organizationId"`
	AuthUserID     string    `gorm:"column:auth_user_id" json:"authUserId"`
	Status         string    `gorm:"column:status" json:"status"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (Membership) TableName() string {
	return "memberships"
}

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		DB: db,
	}
}

type createOrgRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CreateOrgResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (h *Handler) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	authUserID, ok := appauth.UserIDFromContext(r.Context())
	if !ok || strings.TrimSpace(authUserID) == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req createOrgRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Slug = strings.TrimSpace(req.Slug)

	if req.Name == "" || req.Slug == "" {
		http.Error(w, "name and slug are required", http.StatusBadRequest)
		return
	}

	tx := h.DB.Begin()
	if tx.Error != nil {
		http.Error(w, "failed to begin transaction", http.StatusInternalServerError)
		return
	}

	org := Organization{
		Name: req.Name,
		Slug: req.Slug,
	}
	if err := tx.Create(&org).Error; err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	membership := Membership{
		OrganizationID: org.ID,
		AuthUserID:     authUserID,
		Status:         "MEMBER",
	}
	if err := tx.Create(&membership).Error; err != nil {
		tx.Rollback()
		http.Error(w, "failed to create membership", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit().Error; err != nil {
		http.Error(w, "failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(CreateOrgResponse{
		ID:   org.ID,
		Name: org.Name,
		Slug: org.Slug,
	})
}

type ListOrgResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (h *Handler) ListOrganizations(w http.ResponseWriter, r *http.Request) {
	authUserID, ok := appauth.UserIDFromContext(r.Context())
	if !ok || strings.TrimSpace(authUserID) == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var orgs []ListOrgResponse
	err := h.DB.Table("organizations").
		Select("organizations.id, organizations.name, organizations.slug").
		Joins("INNER JOIN memberships ON memberships.organization_id = organizations.id").
		Where("memberships.auth_user_id = ?", authUserID).
		Where("memberships.status IN ?", []string{"VISITOR", "TRIAL", "MEMBER", "ALUMNI"}).
		Order("organizations.id ASC").
		Scan(&orgs).Error
	if err != nil {
		http.Error(w, "failed to list organizations", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(orgs)
}
