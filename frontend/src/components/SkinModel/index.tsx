//import { useRef } from "react";
import { NearestFilter } from "three";
import { ObjectMap } from "@react-three/fiber";
import { useGLTF, useTexture } from "@react-three/drei";
import skinModel from "../../assets/skin_character.gltf";
import skinTexture from "../../assets/character_test.png";

// https://stackoverflow.com/questions/66477582/gltf-file-loaded-partly-black-threejs-react-typescript-how-to-add-color-light
// https://stackoverflow.com/questions/73933300/three-js-react-three-fiber-property-x-does-not-exist-on-type-object3devent

function SkinModel() {
	const gltf = useGLTF(skinModel) as ObjectMap;
	const texture = useTexture(skinTexture);

	// Make texture render pixelated
	texture.magFilter = NearestFilter;

	//const myMesh = useRef();
	return (
		<mesh
			//ref={myMesh}
			receiveShadow
			castShadow
			geometry={(gltf.nodes.Player as THREE.Mesh).geometry}
			material={gltf.materials.Character}
		>
			<meshStandardMaterial map={texture} />
		</mesh>
	);
}

// Model can be preloaded, since its the same for all skins
useGLTF.preload(skinModel);

export default SkinModel;
