package repository

const CountFlagItemQuery = `SELECT COUNT(*) FROM flags;`
const GetFlagsQuery = `SELECT * FROM flags WHERE project_id = $1 AND name ILIKE '%' || $2 || '%' OFFSET $3 LIMIT $4;`
const GetFlagQuery = `SELECT * FROM flags WHERE id = $1;`
const NewFlagQuery = `INSERT INTO flags (project_id,name,active) VALUES ($1,$2,$3) RETURNING *;`
const EditFlagQuery = `UPDATE flags 
						SET name = COALESCE(NULLIF($1,''), name),
							active = COALESCE($2, active)
						WHERE id = $3 RETURNING *;`
const DeleteFlagQuery = `DELETE FROM flags WHERE id = $1;`

const GetFlagContextsCountQuery = `SELECT COUNT(*) FROM flag_contexts WHERE flag_id = $1;`
const GetFlagContextsQuery = `SELECT * FROM flag_contexts WHERE flag_id = $1;`
const GetFlagContextQuery = `SELECT * FROM flag_contexts WHERE id = $1;`
const NewFlagContextQuery = `INSERT INTO flag_contexts (flag_id, name, condition, value) VALUES ($1,$2,$3,$4);`
const EditFlagContextQuery = `UPDATE flag_contexts
								SET name = COALESCE(NULLIF($1,''), name)
									condition = COALESCE(NULLIF($2,''), condition)
									value = COALESCE(NULLIF($3,''), value)
								WHERE id = $4;`
const DeleteFlagContextQuery = `DELETE flag_contexts WHERE id = $1;`