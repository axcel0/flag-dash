import React, { PropsWithChildren } from "react";

const ToggleButton = ({
	children,
	name,
	defaultChecked,
	checked,
	onClick,
	onChange,
	onBlur,
	placeholder,
}: any) => {
	return (
		<label className='inline-flex relative items-center cursor-pointer group text-xl'>
			<input
				type='checkbox'
				className='absolute left-1/2 -translate-x-1/2 peer appearance-none rounded-md'
				name={name}
				defaultChecked={defaultChecked}
				checked={checked}
				placeholder={placeholder}
				onChange={onChange}
				onBlur={onBlur}
				onClick={onClick}
			/>
			<span className='w-16 h-10 flex items-center flex-shrink-0 p-1 bg-gray-300 rounded-full duration-300 ease-in-out peer-checked:bg-green-400 after:w-8 after:h-8 after:bg-white after:rounded-full after:shadow-md after:duration-300 peer-checked:after:translate-x-6 group-hover:after:translate-x-1'></span>
		</label>
	);
};

export default ToggleButton;
