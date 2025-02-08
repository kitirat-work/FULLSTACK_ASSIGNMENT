import * as React from "react";
import { FunctionComponent } from "react";


interface SplashProps {
  
}
 
const Splash: FunctionComponent<SplashProps> = () => {
	React.useEffect(() => {
		document.title = "Splash Screen";
	}, []);

  return ( 
    <div className="wrap">
      <div className="splash">
        <div className="loader"></div>
      </div>
    </div>
   );
}
 
export default Splash;