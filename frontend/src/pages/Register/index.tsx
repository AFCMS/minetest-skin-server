import { Navigate } from "react-router-dom";
import { useRecoilValue } from "recoil";
import { AuthStateIsAuthenticated } from "../../states/auth";

function Register() {
	// Navigate to homepage if already authenticated
	if (useRecoilValue(AuthStateIsAuthenticated)) {
		return <Navigate to="/" />;
	}

	return (
		<div className="flex justify-center align-middle">
			<div className="app-pannel m-10 flex max-w-md flex-col gap-4 p-6">
				<h1 className="panel-text-heading">Register</h1>
				<label className="block">
					<span className="text-slate-800">Username</span>
					<input
						type="text"
						name="username"
						id="username"
						placeholder="GamerPro"
						className="form"
					/>
				</label>
				<label className="block">
					<span className="text-slate-800">Email</span>
					<input
						type="email"
						name="email"
						id="email"
						placeholder="someone@gmail.com"
						className="form"
					/>
				</label>
				<label className="block">
					<span className="text-slate-800">Password</span>
					<input
						type="password"
						name="password"
						id="password"
						placeholder="SuperStrongPassword"
						className="form"
					/>
				</label>

				<div className="mt-2 align-baseline">
					<input
						type="checkbox"
						name="understand"
						id="understand"
						className="form-checkbox"
						defaultChecked={false}
					/>
					<label
						htmlFor="understand"
						className="ml-2 text-sm text-slate-800"
					>
						I aggree to the TOS
					</label>
				</div>
				<button className="button-primary mt-2 w-full">Register</button>
			</div>
		</div>
	);
}

export default Register;
