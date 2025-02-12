import * as React from "react";
import { FunctionComponent } from "react";
import { useUserDispatch, useUserState } from "../../store/UserContext";
import type { User } from "../../model/user";
import UserService from "../../service/UserService";
import { useNavigate } from "react-router-dom";

const USER_ID = "000018b0e1a211ef95a30242ac180002";

interface SplashProps {

}

const Splash: FunctionComponent<SplashProps> = () => {
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
			const res: User = await UserService.GetUserById(USER_ID);
			dispatchUser({ type: 'update', payload: res });
			
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