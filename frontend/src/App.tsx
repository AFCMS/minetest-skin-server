import { BrowserRouter, Routes, Route } from "react-router-dom";

import Header from "./components/Header";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Register from "./pages/Register";
import SearchSkin from "./pages/SearchSkin";
import About from "./pages/About";
import NotFound from "./pages/NotFound";

import "./style/main.css";
import "./style/forms.css";
import "./style/buttons.css";
import "./style/font.css";
import "./style/pannel.css";

function App() {
	return (
		<div className="min-h-screen bg-blue-200">
			<BrowserRouter>
				<Header />
				<Routes>
					<Route path="/" element={<Home />} />
					<Route path="/login" element={<Login />} />
					<Route path="/register" element={<Register />} />
					<Route path="/search" element={<SearchSkin />} />
					<Route path="/about" element={<About />} />
					<Route path="/*" element={<NotFound />} />
				</Routes>
			</BrowserRouter>
		</div>
	);
}

export default App;
