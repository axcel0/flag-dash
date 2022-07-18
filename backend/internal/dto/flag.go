package dto

import "github.com/blastertwist/flag-dash/internal/dao"

// Flag Request Structs
type GetFlagsRequest struct {
	ProjectID 	uint32	`json:"project_id"`
	Filter		string	`json:"filter"`
	Limit		float32	`json:"limit"`
	PageNum		float32	`json:"page_num"`
}

type GetFlagRequest struct {
	ID	uint32	`json:"id"`
}

type NewFlagRequest struct {
	ProjectID	uint32	`json:"project_id"`
	Name		string 	`json:"name"`
	Active		bool	`json:"active"`
}

type EditFlagRequest struct {
	ID		uint32	`json:"id"`
	Name	string	`json:"name"`
	Active	bool	`json:"active"`
}

type DeleteFlagRequest struct {
	ID	uint32	`json:"id"`
}

// Flag Context Request structs
type GetFlagContextsRequest struct {
	FlagID	uint32	`json:"flag_id"`
	Filter	string	`json:"filter"`
	Limit	float32	`json:"limit"`
	PageNum	float32	`json:"page_num"`
}

type GetFlagContextRequest struct {
	ID	uint32	`json:"id"`
}

type NewFlagContextRequest struct {
	FlagID		uint32	`json:"flag_id"`
	Name		string  `json:"name"`
	Condition	string	`json:"condition"`
	Value		string 	`json:"value"`
}

type EditFlagContextRequest struct {
	ID			uint32	`json:"id"`
	Name		string 	`json:"name"`
	Condition	string	`json:"condition"`
	Value		string	`json:"value"`
}

type DeleteFlagContextRequest struct {
	ID	uint32	`json:"id"`
}

// Flag Response Structs
type GetFlagsResponse struct {
	Flags    []*dao.Flag	`json:"flags"`
	Limit		float32			`json:"limit"`
	PageNum		float32			`json:"page_num"`
	MaxPage		float32			`json:"max_page"`
}

type GetFlagResponse struct {
	Status	string		`json:"status"`
	Flag    *dao.Flag	`json:"flag"`
}

type NewFlagResponse struct {
	Status	string		`json:"status"`
	Flag    *dao.Flag	`json:"flag"`
}

type EditFlagResponse struct {
	Status	string		`json:"status"`
	Flag    *dao.Flag	`json:"flag"`
}

type DeleteFlagResponse struct {
	Status	string	`json:"status"`
	Msg    	string	`json:"msg"`
}

// Flag Context Response Structs
type GetFlagContextsResponse struct {
	FlagContexts   []*dao.FlagContext	`json:"flag_contexts"`
	Limit		float32			`json:"limit"`
	PageNum		float32			`json:"page_num"`
	MaxPage		float32			`json:"max_page"`
}

type GetFlagContextResponse struct {
	Status			string				`json:"status"`
	FlagContext     *dao.FlagContext	`json:"flag_context"`
}

type NewFlagContextResponse struct {
	Status			string				`json:"status"`
	FlagContext		*dao.FlagContext	`json:"flag_context"`
}

type EditFlagContextResponse struct {
	Status			string				`json:"status"`
	FlagContext		*dao.FlagContext	`json:"flag_context"`
}

type DeleteFlagContextResponse struct {
	Status		string	`json:"status"`
	Msg			string	`json:"msg"`
}