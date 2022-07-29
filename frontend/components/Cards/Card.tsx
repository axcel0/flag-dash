import { ScriptProps } from "next/script";
import React, { PropsWithChildren } from "react";

interface CardProps {
	childStyle?: string;
	onClick?: any;
}

const Card: React.FC<PropsWithChildren<CardProps>> = ({
	children,
	childStyle,
	onClick,
}) => {
	return (
		<div className={childStyle} onClick={onClick}>
			{children}
		</div>
	);
};

export default Card;
