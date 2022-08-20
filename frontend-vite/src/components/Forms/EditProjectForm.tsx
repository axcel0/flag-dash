import { useFormik } from "formik";

import { FormProps } from ".";

import {
	useEditProjectMutation,
	useGenProjAccessKeyMutation,
} from "../../redux/features/projects/projectsApiSlice";

const EditProjectForm: React.FC<FormProps> = ({
	handleSuccess,
	handleFailed,
	handleCancel,
	data,
}) => {
	const [editProject, { isLoading: isLoading }] = useEditProjectMutation();
	const [generateAccessKey, { isLoading: isLoadingGen }] =
		useGenProjAccessKeyMutation();

	const handleGenAccessKey = async (e: Event) => {
		e.preventDefault();
		try {
			const res = await generateAccessKey({
				id: data?.project?.id,
			}).unwrap();
		} catch (err) {
			console.log(err);
		}
	};

	const editForm = useFormik({
		enableReinitialize: true,
		initialValues: {
			name: data?.project?.name,
		},
		onSubmit: async (values) => {
			try {
				await editProject({
					...values,
					id: data?.project?.id,
				}).unwrap();
				handleSuccess();
			} catch (err) {
				console.log(err);
				handleFailed();
			}
		},
	});

	return (
		<form onSubmit={editForm.handleSubmit}>
			<h3 className='text-3xl my-2'>Edit Project</h3>
			<h3 className='text-lg'>Name </h3>
			<input
				type='text'
				name='name'
				placeholder='Name'
				value={editForm.values.name as any}
				onChange={editForm.handleChange}
				className='rounded border-2 p-2 my-2 w-full'
			/>
			<h3 className='text-lg'>Authorization Key</h3>
			<div className='flex flex-row'>
				<input
					type='text'
					placeholder='Generate key'
					className='rounded border-2 p-2 my-2 w-full'
					defaultValue={data?.project?.access_key}
				/>
				<button
					className='rounded-md px-2 border-2 hover:bg-gray-200 ml-2'
					onClick={(e: any) => handleGenAccessKey(e)}
				>
					Generate
				</button>
			</div>
			<div>
				<button
					className='rounded-md bg-yellow-300 hover:bg-yellow-200 p-3 m-1'
					type='submit'
				>
					<p className='text-lg'>Edit</p>
				</button>
				<button
					onClick={() => handleCancel()}
					className='rounded-md bg-red-300 hover:bg-red-200 p-3 m-1'
				>
					<p className='text-lg'>Cancel</p>
				</button>
			</div>
		</form>
	);
};

export default EditProjectForm;
