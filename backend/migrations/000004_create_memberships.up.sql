CREATE TYPE membership_status AS ENUM (
  'VISITOR',
  'TRIAL',
  'MEMBER',
  'ALUMNI',
  'LEFT',
  'BANNED'
);

CREATE TABLE memberships (
  id BIGSERIAL PRIMARY KEY,
  organization_id BIGINT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
  auth_user_id TEXT NOT NULL,
  status membership_status NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  UNIQUE (organization_id, auth_user_id)
);

CREATE INDEX memberships_auth_user_id_idx ON memberships(auth_user_id);
CREATE INDEX memberships_organization_id_idx ON memberships(organization_id);