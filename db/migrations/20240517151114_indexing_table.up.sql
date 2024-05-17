-- USERS
CREATE INDEX IF NOT EXISTS 
    idx_nip ON users(nip);

CREATE INDEX IF NOT EXISTS
    idx_name ON users(name);

CREATE INDEX IF NOT EXISTS
    idx_id ON users(id);

CREATE INDEX IF NOT EXISTS
    idx_nip_text ON users(cast(nip as text));

CREATE INDEX IF NOT EXISTS
    idx_nip_it ON users(cast(nip as text)) 
    WHERE cast(nip as text) like '615%';

CREATE INDEX IF NOT EXISTS
    idx_nip_nurse ON users(cast(nip as text)) 
    WHERE cast(nip as text) like '303%';

CREATE INDEX IF NOT EXISTS
    idx_created_at_desc ON users(created_at desc);

CREATE INDEX IF NOT EXISTS
    idx_created_at_asc ON users(created_at asc);   


-- MEDICAL_RECORD 