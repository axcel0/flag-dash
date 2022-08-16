import { useFormik } from "formik";
import { useNewProjectMutation } from "../../redux/features/projects/projectsApiSlice";

const CreateProjectForm = ({
	handleSuccess,
	handleFailed,
	handleCancel,
}: any) => {
	const [newPost, { isLoading: isLoadingAdd }] = useNewProjectMutation();
	const newProjForm = useFormik({
		initialValues: {
			name: "",
		},
		onSubmit: async (values) => {
			const res = await newPost(values).unwrap();

			if (res.status === "201") {
				handleSuccess();
			} else {
				handleFailed();
			}
		},
	});

	return (
		<form onSubmit={newProjForm.handleSubmit}>
			<h1 className='text-2xl font-bold'>Create Project</h1>
			<input
				type='text'
				name='name'
				className='rounded border p-3 my-3'
				placeholder='Project Name...'
				onChange={newProjForm.handleChange}
				value={newProjForm.values.name}
			/>

			<div>
				<button
					className='rounded bg-green-300 mr-2 p-2'
					type='submit'
				>
					<p className='text-xl font-medium'>Create</p>
				</button>
				<button
					onClick={() => handleCancel()}
					className='rounded bg-red-200 mx-2 p-2'
				>
					<p className='text-xl font-medium'>Cancel</p>
				</button>
			</div>
		</form>
	);
};

export default CreateProjectForm;
