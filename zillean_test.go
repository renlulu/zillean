package zillean

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	baseURL = "https://api-scilla.zilliqa.com"
)

type testVector struct {
	privateKey string
	publicKey  string
	address    string
}

var testVectors = []testVector{
	{
		privateKey: "b4eb8e8b343e2cce46db4e7571ec1d9654693cca200bc41cc20148355ca62ed9",
		publicKey:  "0314738163b9bb67ad11aa464fe69a1147df263e8970d7dcfd8f993ddd39e81bd9",
		address:    "4baf5fada8e5db92c3d3242618c5b47133ae003c",
	},
	{
		publicKey:  "034ce268ac5a340038d8acebbdd7363611a5b1197916775e32481f5d6b104faf65",
		privateKey: "fd906f4a20a0507813fcb0d8d166cf45c1b17f217b776a003b3c3cc628c7e513",
		address:    "448261915a80cde9bde7c7a791685200d3a0bf4e",
	},
	{
		publicKey:  "02fa7a501f323cc53e070c0a945370368679e7572960ec24d8a0387ef3b50a2285",
		privateKey: "c1fe69c394ee8564d0ea525ff74e5f37df5267cf5c74f2e224757b7d9e594b73",
		address:    "ded02fd979fc2e55c0243bd2f52df022c40ada1e",
	},
	{
		publicKey:  "036f8676e473af20b4cce7d327103de4504a9c00eae7ea03d0a365fb48817ac97f",
		privateKey: "f197920ff83e3bd727560e1e08a4825910e3aa772b40ec4deffe241ee3a51d20",
		address:    "13f06e60297bea6a3c402f6f64c416a6b31e586e",
	},
	{
		publicKey:  "034c39363529c2d4078f72b8c498c4cbc5ba5e10d8666fe06f104a27e0e44242a0",
		privateKey: "3d0c4009da5b4630d292e93fc8df7a55e9b8b31e182a50d113de6e7a7a728b5b",
		address:    "1a90c25307c3cc71958a83fa213a2362d859cf33",
	},
	{
		publicKey:  "026da5bf6c3a3c91c08a3dd7fbce0beaf5d436039c97b07a0f61aa4b9493e46787",
		privateKey: "8a06825735f2f374e5f60f0e08e06ef7169e2ae47532b741700574b9e1603dc0",
		address:    "625abaebd87dae9ab128f3b3ae99688813d9c5df",
	},
	{
		publicKey:  "032a661f9d4ab8dd9818ced2f62f3da14fdd23e68e58d01a4ae186231d7fb609bb",
		privateKey: "715ab4286eccb4e8e3aec0aa8b6fb0e13c105b727e89aae449f87f5df4cfa53f",
		address:    "36ba34097f861191c48c839c9b1a8b5912f583cf",
	},
	{
		publicKey:  "036a94e925bb200dce41a46c9026e023b226dd5b4cb227ce766d60cc8fab218148",
		privateKey: "b9056706ef197125598e7d5ca5ac9c47339b5c894c0656d33113052470baf149",
		address:    "d2453ae76c9a86aae544fca699dbdc5c576aef3a",
	},
	{
		publicKey:  "0247f13639c7597c8ae2467dd29d219c2749260f690d4069930fdeb7866b2bc1fa",
		privateKey: "a6dedd1be3a36945517a6c8a0e1353414debc04c50c9e036361c2a2d95eef947",
		address:    "72220e84947c36118cdbc580454dfaa3b918cd97",
	},
	{
		publicKey:  "02883100f00eab9fb6fd79034bba84235c22ed642470ff6c1db7cf7f782902d9dc",
		privateKey: "57489b87a3e726267ec5150fab6e202694e806ec8fbbe386572cde640bb3ee4b",
		address:    "50f92304c892d94a385ca6ce6cd6950ce9a36839",
	},
	{
		privateKey: "b776d8f068d11b3c3f5b94db0fb30efea05b73ddb9af1bbd5da8182d94245f0b",
		publicKey:  "02cfa555bb63231d167f643f1a23ba66e6ca1458d416ddb9941e95b5fd28df0ac5",
		address:    "171c5f56a5cd412c9b582aa08ad5898caa4a3585",
	},
	{
		privateKey: "24180e6b0c3021aedb8f5a86f75276ee6fc7ff46e67e98e716728326102e91c9",
		publicKey:  "02163fa604c65aebeb7048c5548875c11418d6d106a20a0289d67b59807abdd299",
		address:    "b5c2cdd79c37209c3cb59e04b7c4062a8f5d5271",
	},
	{
		privateKey: "af71626e38926401a6d2fd8fdf91c97f785b8fb2b867e7f8a884351e59ee9aa6",
		publicKey:  "0285c34ff11ea1e06f44d35afe3cc1748b6b122bb06df021a4767db4ef5fbcf1cd",
		address:    "ce24e6f19d436571fd6caa965b5fa93ff66430ef",
	},
	{
		privateKey: "b94289721618d1a8100bea8502fee149bae7313fcbaebf5ad6a867d557e82971",
		publicKey:  "03c7d47ad99dd1db85c9d0b6abe89ffc137f230c991a33890c05436a14974543a2",
		address:    "51c252c01dfdd08d574ab1e67de30633377078db",
	},
	{
		privateKey: "c577711506abf9dbdabd09c5ae66492e77d6450e8d739ce6493728e781150236",
		publicKey:  "0284929ab70d2b39703319a94144802ea9b9a7c8c1a673b10f01738db5b5c40ea8",
		address:    "cac2d4e3f56a58f7e4525dfad5d02127f3c1b94c",
	},
	{
		privateKey: "e9521f5f9f2ac16f4e22681d43f2c1e15163ffe56e5912d57540dd59d71f2877",
		publicKey:  "02db0535d0e4337e5ada6d87b07569308f3f3de19aad15ba1f95108e96f855e171",
		address:    "9a87118ee31e72e1dca6b94c5b59cd95d3d110f4",
	},
	{
		privateKey: "def6ac38d746349abd8a28903c6b98b06156f6f3925d1c32dec9964541fbbfa1",
		publicKey:  "03dc5ca31dff767cb18f92d208fd74e36e800508d3b88bee61680bb7cefa051594",
		address:    "7b01029fe3f9e4a623809ef100856d9f73d7ad96",
	},
	{
		privateKey: "93fe1b32f4f1358d6809982bce53b9e95c083eb889d14b0c2d0f96012074dc6b",
		publicKey:  "02987f1f6a2fe56cd0c2461f2920e3e0d9bd9de96a8febf9246e187e057c6141c4",
		address:    "7a467094ef7886b7dbe157aab1db45abb9c36063",
	},
	{
		privateKey: "efaa6048789fa0282302a5620a2fb0e2c60095a54c42b8312646bbbe1d2ec801",
		publicKey:  "03645a3995ff3309a36d1ec0991fd87ece2ee36e8887f0d5b304accb9152fe7a60",
		address:    "66f26ebf63f9131d36776a3574394375a28e9841",
	},
	{
		privateKey: "ca7de577b3e6968da27088d22e918c039a96af3d4821b7e103560fb6ca1185c4",
		publicKey:  "02f3e03a4ac451a78254fa3d056448db6cbc29d2ef429228c8435141a126c0c07f",
		address:    "77c32e0887e778f9af811503cee8959e8368f86e",
	},
}

func TestZillean_GeneratePrivateKey(t *testing.T) {
	Convey("returns a new private key from random seed", t, func() {
		result, err := NewZillean(baseURL).GeneratePrivateKey()
		So(err, ShouldBeNil)
		So(result, ShouldHaveLength, 64)
		So(result, ShouldHaveSameTypeAs, "string")
	})
}

func TestZillean_VerifyPrivateKey(t *testing.T) {
	Convey("returns true when a valid private key is given ", t, func() {
		for _, vector := range testVectors {
			result := NewZillean(baseURL).VerifyPrivateKey(vector.privateKey)
			So(result, ShouldBeTrue)
		}
	})

	Convey("returns false when a invalid private key is given ", t, func() {
		result := NewZillean(baseURL).VerifyPrivateKey("invalid private key")
		So(result, ShouldBeFalse)
		result = NewZillean(baseURL).VerifyPrivateKey("0000000000000000000000000000000000000000000000000000000000000000")
		So(result, ShouldBeFalse)
		result = NewZillean(baseURL).VerifyPrivateKey("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141")
		So(result, ShouldBeFalse)
	})
}

func TestZillean_GetPublicKeyFromPrivateKey(t *testing.T) {
	Convey("returns a public key from a private key", t, func() {
		for _, vector := range testVectors {
			result, err := NewZillean(baseURL).GetPublicKeyFromPrivateKey(vector.privateKey)
			So(err, ShouldBeNil)
			So(result, ShouldEqual, vector.publicKey)
		}
	})
}

func TestZillean_IsPublicKey(t *testing.T) {
	Convey("returns true when a valid public key is given ", t, func() {
		for _, vector := range testVectors {
			result := NewZillean(baseURL).IsPublicKey(vector.publicKey)
			So(result, ShouldBeTrue)
		}
	})

	Convey("returns false when a invalid public key is given ", t, func() {
		result := NewZillean(baseURL).IsPublicKey("invalid public key")
		So(result, ShouldBeFalse)
	})
}

func TestZillean_GetAddressFromPrivateKey(t *testing.T) {
	Convey("returns an address from a private key", t, func() {
		for _, vector := range testVectors {
			result, err := NewZillean(baseURL).GetAddressFromPrivateKey(vector.privateKey)
			So(err, ShouldBeNil)
			So(result, ShouldEqual, vector.address)
		}
	})
}

func TestZillean_GetAddressFromPublicKey(t *testing.T) {
	Convey("returns an address from a public key", t, func() {
		for _, vector := range testVectors {
			result, err := NewZillean(baseURL).GetAddressFromPublicKey(vector.publicKey)
			So(err, ShouldBeNil)
			So(result, ShouldEqual, vector.address)
		}
	})
}

func TestZillean_IsAddress(t *testing.T) {
	Convey("returns true when a valid address is given ", t, func() {
		for _, vector := range testVectors {
			result := NewZillean(baseURL).IsAddress(vector.address)
			So(result, ShouldBeTrue)
		}
	})

	Convey("returns false when a invalid address is given ", t, func() {
		result := NewZillean(baseURL).IsAddress("invalid address")
		So(result, ShouldBeFalse)
	})
}
