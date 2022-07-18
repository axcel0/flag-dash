package repository

const GetProjectCountQuery = `SELECT COUNT(*) FROM projects;`
const GetProjectsQuery = `SELECT * FROM projects WHERE name ILIKE '%' || $1 || '%' OFFSET $2 LIMIT $3;`
const GetProjectQuery = `SELECT * FROM projects WHERE id = $1;`
const NewProjectQuery = `INSERT INTO projects (name) VALUES ($1) RETURNING *;`
const EditProjectQuery = `UPDATE projects SET name = COALESCE(NULLIF($1, ''), name), update_time = now() WHERE id = $2 RETURNING *;`
const DeleteProjectQuery = `DELETE FROM projects WHERE id = $1;`