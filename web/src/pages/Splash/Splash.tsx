import * as React from "react";
import { FunctionComponent } from "react";
import { useNavigate } from "react-router-dom";
import UserService from "../../service/UserService";
import { useUserDispatch, useUserState } from "../../store/UserContext";


interface SplashProps {
	
}

const Splash: FunctionComponent<SplashProps> = () => {
	const queryParams = new URLSearchParams(window.location.search);
	const userId = queryParams.get('userId');
	const dispatchUser = useUserDispatch();
	const { user } = useUserState();
	let navigate = useNavigate();

	React.useEffect(() => {
		try {
			document.title = "Splash Screen";
			if (!user) {
				initUser();
			} else {
				navigate('/pin');
			}
		} catch (error) {
			console.error(error);
		}
	}, [user]);

	async function initUser() {
		try {
			const res = await UserService.GetUserById(userId);
			dispatchUser({ type: 'update', payload: res.data });
			
		} catch (error) {
			console.error(error);
		}
	}
	return React.useMemo(() => (
		<div className="splash">
			<div className="loader"></div>
		</div>
	), []);
};

export default Splash;