package main

import (
	tfe "github.com/hashicorp/go-tfe"
)

func getStructs() map[string]interface{} {
	s := make(map[string]interface{})

	s["workspace_create_options"] = &tfe.WorkspaceCreateOptions{}
	s["workspace_list_options"] = &tfe.WorkspaceListOptions{}
	s["workspace_lock_options"] = &tfe.WorkspaceLockOptions{}
	s["workspace_update_options"] = &tfe.WorkspaceUpdateOptions{}

	return s
}
