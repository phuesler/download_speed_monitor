CREATE TABLE statistics (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  started_at INTEGER,
  duration_ms INTEGER,
  md5_source VARCHAR(40),
  md5_target VARCHAR(40),
  file_size_bytes INTEGER,
  error_message VARCHAR(255)
);
