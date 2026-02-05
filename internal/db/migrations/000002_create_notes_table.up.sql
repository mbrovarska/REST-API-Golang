CREATE TABLE IF NOT EXISTS notes (
    id         BIGSERIAL PRIMARY KEY,
	title      VARCHAR(255) NOT NULL,
	content    TEXT,
	user_id    BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_notes_user_id ON notes(user_id);
CREATE INDEX idx_notes_created_at ON notes(created_at DESC);