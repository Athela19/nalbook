CREATE TABLE IF NOT EXISTS schema_migrations (
  version TEXT PRIMARY KEY,
  applied_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS companies (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  code TEXT NOT NULL,
  name TEXT NOT NULL,
  base_currency CHAR(3) NOT NULL DEFAULT 'IDR',
  coa_template TEXT NOT NULL DEFAULT 'jasa',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE (code)
);

CREATE TABLE IF NOT EXISTS period_locks (
  company_id UUID NOT NULL REFERENCES companies(id) ON DELETE RESTRICT,
  period_start DATE NOT NULL,
  period_end DATE NOT NULL,
  is_closed BOOLEAN NOT NULL DEFAULT FALSE,
  closed_by UUID NULL,
  closed_at TIMESTAMPTZ NULL,
  reason TEXT NULL,
  PRIMARY KEY (company_id, period_start, period_end)
);

INSERT INTO schema_migrations (version) VALUES ('0001_init')
ON CONFLICT (version) DO NOTHING;
