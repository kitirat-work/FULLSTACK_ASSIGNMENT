import * as React from "react";
import { FunctionComponent } from "react";
import { useUserDispatch, useUserState } from "../../store/UserContext";
import type { User } from "../../model/user";
import UserService from "../../service/UserService";
import { useNavigate, useParams } from "react-router-dom";


interface SplashProps {
	
}

const Splash: FunctionComponent<SplashProps> = () => {
	const { userId } = useParams<{ userId: string }>();
	const USER_ID = userId || "defaultUserId";
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
			const res = await UserService.GetUserById(USER_ID);
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