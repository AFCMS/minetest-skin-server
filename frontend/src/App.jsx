import { BrowserRouter, Routes, Route } from "react-router-dom";

import Header from "./components/Header";
import Home from "./pages/Home";
import Login from "./pages/Login";
import About from "./pages/About";
import NotFound from "./pages/NotFound";

import "./style/buttons.css";

function App() {
	return (
		<div className="bg-blue-200 min-h-screen">
			<BrowserRouter>
				<Header />
				<Routes>
					<Route path="/" element={<Home />} />
					<Route path="/login" element={<Login />} />
					<Route path="/about" element={<About />} />
					<Route path="/*" element={<NotFound />} />
				</Routes>
			</BrowserRouter>
		</div>
	);
}

export default App;
