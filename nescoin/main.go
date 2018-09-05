package main

import (
	"./ecdsacrypto"
)

func main() {
	// cryptoutils.GeneratePair()
	// cryptoutils.SignMessage("Hello World!")
	// 	const publicKey = `
	// -----BEGIN PUBLIC KEY-----
	// MIICCgKCAgEAu6hcNFLJpYPy5ljZW6Bu3ec6xzjsG20rf+xAk6/Thb4xMEalNKHc
	// JfPP9jbme+Fm/0rX6smttpa4QfzxFQ1RgrUSWT9hkH3Q5Hk2Yogi1Xizp4jRkW/I
	// Lbb5AWx+p2uExvELawfJcMdOCkgyN7J2djySWGk1NzFqeJkBVrNq7D5ouC7g/hrE
	// BCqHdXfJxlk1bqNWXByqXUVwqlakPAxqzUOE+IazsnQL/sZ3yHeWsMuzpII8qHKY
	// KijL0MU6WtlVguNf77OG4ZkLwZ2PeQY1ewYKIfr2DDij1LrivG9kejikvdWFCHq7
	// 87RFraBv8XGVvI8/loZ2EcRqaDFMVHc8OQyKqh5XBfGV5BmNP0dsva0J5LYhlLeo
	// vuHtSRofVsU/WXDP7VvOwLLGM0m6Tl20Ok8W1JBP+F3ts5rPFuLdU3eViPrlBXFK
	// i0UZwZOsOKtVEIPjJXUG8yImCTCuPrW5ea7lM+WgsNuxjtcHVccbiXNeVGuy+YTX
	// PoADlAVvE9xGoydDRIylCAuKSqRR1/WOFz64pFmX086PkxgHWYOCRNm63j324Kgk
	// i3wIE4mjVUQlXxe9OwOTfxjmdIRpa+UALaCO8gJEGexkVygghAWXnuyasgtIeO91
	// YdlMD6alM4kGAs3O2JdW9uFc3xkAlyLnEZkaE3MXaSOQChiqG2orlOkCAwEAAQ==
	// -----END PUBLIC KEY-----`
	// 	message := "Hello World!"
	// 	signature := "07fac6ef3985b8d280927cbbcb5f42fe0fc4932380b139d6846dd1443f5cd9d5202f3f2b12d080505fff9a202bcc4f739d43b11d69ecdd40be970bd14e10d3d19011599b34689913645ccee54f8c47a57992213bbf78997812ae3def808b3cd9f8e9c05864ea0d342914e5d1fc754d39fb1c990fa0c3464aa91cb6d9ac182d5810518f76d57203caa9cb8069ba9b4eef80ec36f51a5887331ee3f30ab5f71088121237dee84e21f703c12c8cc29cc3bbefafcb09a8797d647c1e8372b73dfa451eda80694009fba170c7b0cc3f1916ded4544c240e943b1902ba726d8cf918c759a9fe14decd03f8d9a4d6133e439f656dcd2735ff2d75b86ce802a3f2a0f682c36816ce68c0a8459d4ede4662f6735127a4b6e1b0591276ca315ae2e85c840095aa460b689b7c8b40cf29ed2520d9c91c6429a64793a4d8158b9e767e72b3790cbc66eeb4c567d7d19f84274896663399f15233f009e6aa4c89711cd02d6b0c661ce4faa3448960adf6dcc439a7f81506616f3a98113e34b6cf789d73702371b2a2c1722b3d92a575c6c8d0cbe52b10be41fa6d504e57b4bc9b6faca19d165233507ee30b060ef5790e5d8f23773723ca6fb2208861f9efe23ba3f16e04c68cf84bb2d2e95b0e2801021f32ed08fa017d841067cea3a0170010216eea1d3a2d172e2182ad3c94da43a57b6dafeb8cd5cce5848c832a2cc53d2304d5a1959d81"
	// 	cryptoutils.ValidateSignature([]byte(publicKey), signature, message)
	//ecdsacrypto.KeyGen()
	ecdsacrypto.Sign("Hello World!")
	// 	r := big.NewInt(0)
	// 	r.SetBytes([]byte("35916928212251723310950872460988008332733186530634149852157469200804733997700"))
	// 	s := big.NewInt(0)
	// 	s.SetBytes([]byte("114445191554386117141866967283627901438174717970749003808786873860689020514003"))
	// 	const publicKey = `
	// -----BEGIN PUBLIC KEY-----
	// MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEUJcRfEQ1cETgQ7wEtk/6v3Drdy40
	// mcbfJ7OEHLSgwn0AbOjukBP/oboIfJgbbKPYp6Y1cZCoogJJdnFpX1mJcA==
	// -----END PUBLIC KEY-----`
	//ecdsacrypto.Verify("Hello World!", r, s, []byte(publicKey))
}
