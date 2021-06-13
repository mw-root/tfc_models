package main

import (
	tfe "github.com/hashicorp/go-tfe"
)

func getStructs() map[string]interface{} {
	s := make(map[string]interface{})

	s["organizatoin_create_options"] = &tfe.OrganizationCreateOptions{}
	s["organization_list_options"] = &tfe.OrganizationListOptions{}
	s["organization_update_options"] = &tfe.OrganizationUpdateOptions{}
	s["run_queue_options"] = &tfe.RunQueueOptions{}
	//s[""] = &tfe.
	//s[""] = &tfe.
	//s[""] = &tfe.
	//s[""] = &tfe.
	//s[""] = &tfe.
	//s[""] = &tfe.
	s["organization_membership_create_options"] = &tfe.OrganizationMembershipCreateOptions{}
	s["organization_membership_list_options"] = &tfe.OrganizationMembershipListOptions{}
	s["organization_membership_read_options"] = &tfe.OrganizationMembershipReadOptions{}
	s["plan_export_create_options"] = &tfe.PlanExportCreateOptions{}
	s["policy_list_options"] = &tfe.PolicyListOptions{}
	s["policy_create_options"] = &tfe.PolicyCreateOptions{}
	s["policy_update_options"] = &tfe.PolicyUpdateOptions{}
	s["policy_check_list_options"] = &tfe.PolicyCheckListOptions{}
	s["policy_set_add_policies_options"] = &tfe.PolicySetAddPoliciesOptions{}
	s["policy_set_add_workspaces_options"] = &tfe.PolicySetAddWorkspacesOptions{}
	s["policy_set_create_options"] = &tfe.PolicySetCreateOptions{}
	s["policy_set_list_options"] = &tfe.PolicySetListOptions{}
	s["policy_set_parameter_update_options"] = &tfe.PolicySetParameterUpdateOptions{}
	s["policy_set_parameter_create_options"] = &tfe.PolicySetParameterCreateOptions{}
	s["policy_set_parameter_list_options"] = &tfe.PolicySetParameterListOptions{}
	s["policy_set_remove_policies_options"] = &tfe.PolicySetRemovePoliciesOptions{}
	s["policy_set_remove_workspace_options"] = &tfe.PolicySetRemoveWorkspacesOptions{}
	s["policy_set_update_options"] = &tfe.PolicySetUpdateOptions{}
	//s["policy_set_version"] = &tfe.PolicySetVersion{}
	s["registry_module_create_options"] = &tfe.RegistryModuleCreateOptions{}
	s["registry_module_create_version_options"] = &tfe.RegistryModuleCreateVersionOptions{}
	s["registry_module_create_with_vcs_connection_options"] = &tfe.RegistryModuleCreateWithVCSConnectionOptions{}
	s["registry_module_verison"] = &tfe.RegistryModuleVersion{}
	s["run_apply_options"] = &tfe.RunApplyOptions{}
	s["run_cancel_options"] = &tfe.RunCancelOptions{}
	s["run_create_options"] = &tfe.RunCreateOptions{}
	s["run_discard_options"] = &tfe.RunDiscardOptions{}
	s["run_force_cancel_options"] = &tfe.RunForceCancelOptions{}
	s["run_list_options"] = &tfe.RunListOptions{}
	s["run_read_options"] = &tfe.RunReadOptions{}
	s["run_trigger_create_options"] = &tfe.RunTriggerCreateOptions{}
	s["run_trigger_list_options"] = &tfe.RunTriggerListOptions{}
	s["ssh_key_create_options"] = &tfe.SSHKeyCreateOptions{}
	s["ssh_key_list_options"] = &tfe.SSHKeyListOptions{}
	s["ssh_key_update_options"] = &tfe.SSHKeyUpdateOptions{}
	s["state_version_create_options"] = &tfe.StateVersionCreateOptions{}
	s["state_version_current_options"] = &tfe.StateVersionCurrentOptions{}
	s["state_version_list_options"] = &tfe.StateVersionListOptions{}
	s["state_version_read_options"] = &tfe.StateVersionReadOptions{}
	s["team_member_add_options"] = &tfe.TeamMemberAddOptions{}
	s["team_member_remove_options"] = &tfe.TeamMemberRemoveOptions{}
	s["user_update_options"] = &tfe.UserUpdateOptions{}
	s["user_token_generate_options"] = &tfe.UserTokenGenerateOptions{}
	s["variable_create_options"] = &tfe.VariableCreateOptions{}
	s["variable_list_options"] = &tfe.VariableListOptions{}
	s["variable_update_options"] = &tfe.VariableUpdateOptions{}
	s["workspace_assign_sshkey_options"] = &tfe.WorkspaceAssignSSHKeyOptions{}
	s["workspace_add_remote_state_consumers_options"] = &tfe.WorkspaceAddRemoteStateConsumersOptions{}
	s["workspace_create_options"] = &tfe.WorkspaceCreateOptions{}
	s["workspace_list_options"] = &tfe.WorkspaceListOptions{}
	s["workspace_lock_options"] = &tfe.WorkspaceLockOptions{}
	s["workspace_remove_remote_state_consumers_options"] = &tfe.WorkspaceRemoveRemoteStateConsumersOptions{}
	s["workspace_update_options"] = &tfe.WorkspaceUpdateOptions{}
	s["workspace_update_remote_state_consumers_options"] = &tfe.WorkspaceUpdateRemoteStateConsumersOptions{}

	return s
}
