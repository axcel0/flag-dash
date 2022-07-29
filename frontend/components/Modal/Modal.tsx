import React, { PropsWithChildren } from "react";

interface ModalProps {
	visible: boolean;
	onClose: any;
	childStyle?: string;
}
const Modal: React.FC<PropsWithChildren<ModalProps>> = ({
	children,
	childStyle,
	visible,
	onClose,
}) => {
	if (!visible) return null;
	return (
		<div
			id='container'
			onClick={(e) =>
				(e.target as HTMLDivElement).id === "container"
					? onClose()
					: null
			}
			className='fixed inset-0 fixed z-10 bg-black bg-opacity-30 backdorp-blur-sm w-screen h-screen overflow-auto flex justify-center items-center'
		>
			<div id='childContainer' className={childStyle + ""}>
				{children}
			</div>
		</div>
	);
};

export default Modal;
