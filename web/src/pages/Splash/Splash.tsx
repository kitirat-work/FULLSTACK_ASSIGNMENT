import * as React from "react";
import { FunctionComponent } from "react";
import { useUserDispatch, useUserState } from "../../store/UserContext";
import type { User } from "../../model/user";
import UserService from "../../service/UserService";

const USER_ID = "000018b0e1a211ef95a30242ac180002";

interface SplashProps {

}

const Splash: FunctionComponent<SplashProps> = () => {
	const { user } = useUserState();
	const dispatchUser = useUserDispatch();

	React.useEffect(() => {
		try {
			document.title = "Splash Screen";
			if (!user) {
				initUser();
			}
			// window.location.href = "/pin";
		} catch (error) {
			console.error(error);
		}
	}, []);

	async function initUser() {
		try {
			const res: User = await UserService.getUserById(USER_ID);
			dispatchUser({ type: 'update', payload: res });
		} catch (error) {
			console.error(error);
		}
	}
	return (
		<div className="splash">
			<div className="loader"></div>
		</div>
	);
};

export default Splash;