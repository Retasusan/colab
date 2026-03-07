CREATE TABLE jwks (
  "id" TEXT NOT NULL PRIMARY KEY,
  "publicKey" TEXT NOT NULL,
  "privateKey" TEXT NOT NULL,
  "createdAt" timestamptz NOT NULL,
  "expiresAt" timestamptz
);

