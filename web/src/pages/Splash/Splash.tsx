import * as React from "react";
import { FunctionComponent } from "react";


interface SplashProps {

}

const Splash: FunctionComponent<SplashProps> = () => {
	React.useEffect(() => {
		document.title = "Splash Screen";
	}, []);
	// get user id from local storage
	// get user info from api
	// if user info is null end!!
	// redirect to pin screen
	return (
		<div className="splash">
			<div className="loader"></div>
		</div>
	);
};

export default Splash;