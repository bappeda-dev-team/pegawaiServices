CREATE TABLE tb_pegawai (
    id BIGSERIAL PRIMARY KEY,
    nama TEXT NOT NULL,
    nip VARCHAR(255) NOT NULL UNIQUE,
    kode_opd VARCHAR(255) NOT NULL,
    status_pegawai VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_tb_pegawai_updated_at
BEFORE UPDATE ON tb_pegawai
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();