export { default as CreateProjectForm } from "./CreateProjectForm";
export { default as CreateFlagForm } from "./CreateFlagForm";
export { default as EditFlagForm } from "./EditFlagForm";
export { default as EditProjectForm } from "./EditProjectForm";

export interface FormProps {
	handleSuccess: any;
	handleFailed: any;
	handleCancel: any;
	data?: any;
}
