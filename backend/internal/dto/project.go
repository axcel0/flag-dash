package dto

import "github.com/blastertwist/flag-dash/internal/dao"

type NewProjectRequest struct {
	Name	string	`json:"name"`
}

type GetProjectsRequest struct {
	Filter 	string `query:"filter"`
	Limit 	float32 `query:"limit"`
	PageNum	float32 `query:"page_num"`
}

type GetProjectRequest struct {
	ID	float32 `params:"id" json:"id"`
}

type EditProjectRequest struct {
	ID		uint32	
	Name	string  `json:"name"`
}

type DeleteProjectRequest struct {
	ID	uint32	`json:"id"`
}

type GetProjectsResponse struct {
	Projects	[]*dao.Project	`json:"projects"`
	Limit		float32			`json:"limit"`
	PageNum		float32			`json:"page_num"`
	MaxPage		float32			`json:"max_page"`
}

type GetProjectResponse struct {
	Status		string			`json:"status"`
	Project		*dao.Project	`json:"project"`
}

type NewProjectReponse struct {
	Status		string			`json:"status"`
	Project     *dao.Project	`json:"project"`
}

type EditProjectResponse struct {
	Status		string			`json:"status"`
	Project     *dao.Project	`json:"project"`
}

type DeleteProjectResponse struct {
	Status		string		`json:"status"`
	Msg			string		`json:"msg"`
}