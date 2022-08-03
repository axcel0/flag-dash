import React, { useState, useEffect } from "react";

import { Navigate, useLocation, Outlet, useNavigate } from "react-router-dom";

const RequireAuth = () => {
	const location = useLocation();
	const navigate = useNavigate();

	return <Outlet />;
};

export default RequireAuth;
