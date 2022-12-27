import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

function Login() {
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");

	const [canSubmit, setCanSubmit] = useState(false);

	useEffect(() => {
		if (email === "" && password === "") {
			setCanSubmit(false);
		} else {
			setCanSubmit(true);
		}
	}, [email, password, setCanSubmit]);

	const navigate = useNavigate();

	return (
		<div className="justify-center align-middle flex">
			<div className="flex flex-col max-w-md p-6 gap-4 m-10 app-pannel">
				<h1 className="panel-text-heading">Login</h1>
				<label className="block">
					<span className="text-slate-800">Email</span>
					<input
						type="email"
						name="email"
						id="email"
						placeholder="Ex: someone@gmail.com"
						className="form"
						value={email}
						onChange={(e) => {
							setEmail(e.target.value);
						}}
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
						value={password}
						onChange={(e) => {
							setPassword(e.target.value);
						}}
					/>
				</label>
				<div className="flex gap-4 mt-4">
					<button
						className="button-secondary w-full"
						onClick={(e) => {
							navigate("/register");
						}}
					>
						Register
					</button>
					<button
						className="button-primary w-full disabled:bg-black disabled:cursor-not-allowed"
						disabled={!canSubmit}
					>
						Login
					</button>
				</div>
			</div>
		</div>
	);
}

export default Login;
