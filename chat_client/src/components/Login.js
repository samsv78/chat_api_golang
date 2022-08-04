import React from 'react'

const Login = (props) => {
    return (
        <>
            <div className="limiter">
                <div className="container-login100">
                    <div className="wrap-login100">
                        <span className="login100-form-title p-b-26">Welcome</span>
                        <span className="login100-form-title p-b-48">
                            <i className="zmdi zmdi-font" />
                        </span>
                        <div
                            className="wrap-input100 validate-input"
                            data-validate="Valid email is: a@b.c"
                        >
                            <input className="input100" type="text" name="email" id='email' />
                            <span className="focus-input100" data-placeholder="" />
                        </div>
                        <div
                            className="wrap-input100 validate-input"
                            data-validate="Enter password"
                        >
                            <span className="btn-show-pass">
                                <i className="zmdi zmdi-eye" />
                            </span>
                            <input className="input100" type="password" name="pass" id='password' />
                            <span className="focus-input100" data-placeholder="" />
                        </div>
                        <div className="container-login100-form-btn">
                            <div className="wrap-login100-form-btn">
                                <div className="login100-form-bgbtn" />
                                <button onClick={props.login} className="login100-form-btn">Login</button>
                            </div>
                        </div>
                        {/* <div className="text-center p-t-115">
                                <span className="txt1">Don’t have an account?</span>
                                <a className="txt2" href="#">
                                    Sign Up
                                </a>
                            </div> */}
                    </div>
                </div>
            </div>
            <div id="dropDownSelect1" />

        </>

    )
}

export default Login