package app

import (
	st "github.com/NautiloosGo/url/internal/storage"
	"testing"
)

type ask struct {
	str []string
}

type ans struct {
	str []string
	ok  []bool
}

var tCat = st.Catalog{List: make([]st.Request, 0)}
var tReq = st.Request{
	Id:   "",
	Url:  "",
	Surl: "",
}

func TestFindSurlTrue(t *testing.T) {
	//Arrange
	lista := []string{
		"0123456789",
		"ABCDEFGHIJ",
		"abcdefghij",
		"__________",
		"1111111112",
		"1111111113",
		"1111111114",
	}

	listb := []string{
		"1234567890",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghijklmnopqrstuvwxyz",
		`!"#$%&'()*+,-./:;<=>?@"`,
		`[\]^_`,
		`{|}~¡¢£¤¥¦§¨©ª«¬­®¯°±²³´µ¶·¸¹º»¼½¾¿`,
		`ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷øùúûüýþÿĀāĂăĄąĆćĈĉĊċČčĎďĐđĒēĔĕĖėĘęĚěĜĝĞğĠġĢģĤĥĦħĨĩĪīĬĭĮįİıĲĳĴĵĶķĸĹĺĻļĽľĿŀŁłŃńŅņŇňŉŊŋŌōŎŏŐőŒœŔŕŖŗŘřŚśŜŝŞşŠšŢţŤťŦŧŨũŪūŬŭŮůŰűŲųŴŵŶŷŸŹźŻżŽžſƀƁƂƃƄƅƆƇƈƉƊƋƌƍƎƏƐƑƒƓƔƕƖƗƘƙƚƛƜƝƞƟƠơƢƣƤƥƦƧƨƩƪƫƬƭƮƯưƱƲƳƴƵƶƷƸƹƺƻƼƽƾƿǀǁǂǃǄǅǆǇǈǉǊǋǌǍǎǏǐǑǒǓǔǕǖǗǘǙǚǛǜǝǞǟǠǡǢǣǤǥǦǧǨǩǪǫǬǭǮǯǰǱǲǳǴǵǶǷǸǹǺǻǼǽǾǿȀȁȂȃȄȅȆȇȈȉȊȋȌȍȎȏȐȑȒȓȔȕȖȗȘșȚțȜȝȞȟȠȡȢȣȤȥȦȧȨȩȪȫȬȭȮȯȰȱȲȳȴȵȶȷȸȹȺȻȼȽȾȿɀɁɂɃɄɅɆɇɈɉɊɋɌɍɎɏɐɑɒɓɔɕɖɗɘəɚɛɜɝɞɟɠɡɢɣɤɥɦɧɨɩɪɫɬɭɮɯɰɱɲɳɴɵɶɷɸɹɺɻɼɽɾɿʀʁʂʃʄʅʆʇʈʉʊʋʌʍʎʏʐʑʒʓʔʕʖʗʘʙʚʛʜʝʞʟʠʡʢʣʤʥʦʧʨʩʪʫʬʭʮ`,
	}

	goalOK := true

	// Act
	listans := make([]string, len(lista))
	ok := make([]bool, len(lista))
	tCatalog := createCat()
	for i, u := range lista {
		listans[i], ok[i] = FindSurl(tCatalog, u)
	}
	// Assert
	for j, surl := range listans {
		if listb[j] != surl || ok[j] != goalOK {
			t.Fatalf("%q: \nrequest: %s \nwant %s and %t \nresult: %s and %t", t.Name(), lista[j], listb[j], goalOK, listans[j], ok[j])
		}
	}
}

func TestFindSurlFalse(t *testing.T) {
	//Arrange

	lista := []string{
		"4043456789",
		"ZZZDEFGHIJ",
		"zzzdefghij",
		"_____!____",
		"2111111112",
		"2111111113",
		"2111111114",
	}

	goalOK := false

	// Act
	listans := make([]string, len(lista))
	ok := make([]bool, len(lista))
	tCatalog := createCat()
	for i, u := range lista {
		listans[i], ok[i] = FindSurl(tCatalog, u)
	}
	// Assert
	for j, _ := range listans {
		if ok[j] != goalOK {
			t.Fatalf("%q: \nrequest: %s \nwant not found and ok=%t \nbut found %s and %t", t.Name(), lista[j], goalOK, listans[j], ok[j])
		}
	}
}

func TestFindUrlTrue(t *testing.T) {
	//Arrange

	lista := []string{
		"1234567890",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghijklmnopqrstuvwxyz",
		`!"#$%&'()*+,-./:;<=>?@"`,
		`[\]^_`,
		`{|}~¡¢£¤¥¦§¨©ª«¬­®¯°±²³´µ¶·¸¹º»¼½¾¿`,
		`ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷øùúûüýþÿĀāĂăĄąĆćĈĉĊċČčĎďĐđĒēĔĕĖėĘęĚěĜĝĞğĠġĢģĤĥĦħĨĩĪīĬĭĮįİıĲĳĴĵĶķĸĹĺĻļĽľĿŀŁłŃńŅņŇňŉŊŋŌōŎŏŐőŒœŔŕŖŗŘřŚśŜŝŞşŠšŢţŤťŦŧŨũŪūŬŭŮůŰűŲųŴŵŶŷŸŹźŻżŽžſƀƁƂƃƄƅƆƇƈƉƊƋƌƍƎƏƐƑƒƓƔƕƖƗƘƙƚƛƜƝƞƟƠơƢƣƤƥƦƧƨƩƪƫƬƭƮƯưƱƲƳƴƵƶƷƸƹƺƻƼƽƾƿǀǁǂǃǄǅǆǇǈǉǊǋǌǍǎǏǐǑǒǓǔǕǖǗǘǙǚǛǜǝǞǟǠǡǢǣǤǥǦǧǨǩǪǫǬǭǮǯǰǱǲǳǴǵǶǷǸǹǺǻǼǽǾǿȀȁȂȃȄȅȆȇȈȉȊȋȌȍȎȏȐȑȒȓȔȕȖȗȘșȚțȜȝȞȟȠȡȢȣȤȥȦȧȨȩȪȫȬȭȮȯȰȱȲȳȴȵȶȷȸȹȺȻȼȽȾȿɀɁɂɃɄɅɆɇɈɉɊɋɌɍɎɏɐɑɒɓɔɕɖɗɘəɚɛɜɝɞɟɠɡɢɣɤɥɦɧɨɩɪɫɬɭɮɯɰɱɲɳɴɵɶɷɸɹɺɻɼɽɾɿʀʁʂʃʄʅʆʇʈʉʊʋʌʍʎʏʐʑʒʓʔʕʖʗʘʙʚʛʜʝʞʟʠʡʢʣʤʥʦʧʨʩʪʫʬʭʮ`,
	}

	listb := []string{
		"0123456789",
		"ABCDEFGHIJ",
		"abcdefghij",
		"__________",
		"1111111112",
		"1111111113",
		"1111111114",
	}

	goalOK := true

	// Act
	listans := make([]string, len(lista))
	ok := make([]bool, len(lista))
	tCatalog := createCat()
	for i, u := range lista {
		listans[i], ok[i] = FindUrl(tCatalog, u)
	}
	// Assert
	for j, surl := range listans {
		if listb[j] != surl || ok[j] != goalOK {
			t.Fatalf("%q: \nrequest: %s \nwant %s and %t \nresult: %s and %t", t.Name(), lista[j], listb[j], goalOK, listans[j], ok[j])
		}
	}
}

func TestFindUrlFalse(t *testing.T) {
	//Arrange

	lista := []string{
		"0000000000",
		"AAAAAAAAAAA",
		"aaaaaaaaaaaa",
		`!!!!!!!"#$%&'()*+,-./:;<=>?@"`,
		`___________[\]^_`,
		`|||||||||||{|}~¡¢£¤¥¦§¨©ª«¬­®¯°±²³´µ¶·¸¹º»¼½¾¿`,
		`ÀÀÀÀÀÀÀÀÀÀÀÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷øùúûüýþÿĀāĂăĄąĆćĈĉĊċČčĎďĐđĒēĔĕĖėĘęĚěĜĝĞğĠġĢģĤĥĦħĨĩĪīĬĭĮįİıĲĳĴĵĶķĸĹĺĻļĽľĿŀŁłŃńŅņŇňŉŊŋŌōŎŏŐőŒœŔŕŖŗŘřŚśŜŝŞşŠšŢţŤťŦŧŨũŪūŬŭŮůŰűŲųŴŵŶŷŸŹźŻżŽžſƀƁƂƃƄƅƆƇƈƉƊƋƌƍƎƏƐƑƒƓƔƕƖƗƘƙƚƛƜƝƞƟƠơƢƣƤƥƦƧƨƩƪƫƬƭƮƯưƱƲƳƴƵƶƷƸƹƺƻƼƽƾƿǀǁǂǃǄǅǆǇǈǉǊǋǌǍǎǏǐǑǒǓǔǕǖǗǘǙǚǛǜǝǞǟǠǡǢǣǤǥǦǧǨǩǪǫǬǭǮǯǰǱǲǳǴǵǶǷǸǹǺǻǼǽǾǿȀȁȂȃȄȅȆȇȈȉȊȋȌȍȎȏȐȑȒȓȔȕȖȗȘșȚțȜȝȞȟȠȡȢȣȤȥȦȧȨȩȪȫȬȭȮȯȰȱȲȳȴȵȶȷȸȹȺȻȼȽȾȿɀɁɂɃɄɅɆɇɈɉɊɋɌɍɎɏɐɑɒɓɔɕɖɗɘəɚɛɜɝɞɟɠɡɢɣɤɥɦɧɨɩɪɫɬɭɮɯɰɱɲɳɴɵɶɷɸɹɺɻɼɽɾɿʀʁʂʃʄʅʆʇʈʉʊʋʌʍʎʏʐʑʒʓʔʕʖʗʘʙʚʛʜʝʞʟʠʡʢʣʤʥʦʧʨʩʪʫʬʭʮ`,
	}

	goalOK := false

	// Act
	listans := make([]string, len(lista))
	ok := make([]bool, len(lista))
	tCatalog := createCat()
	for i, u := range lista {
		listans[i], ok[i] = FindSurl(tCatalog, u)
	}
	// Assert
	for j, _ := range listans {
		if ok[j] != goalOK {
			t.Fatalf("%q: \nrequest: %s \nwant not found and ok=%t \nbut found %s and %t", t.Name(), lista[j], goalOK, listans[j], ok[j])
		}
	}
}

func createCat() st.Catalog {
	tReq.Url = "1234567890"
	tReq.Surl = "0123456789"
	tCat.List = append(tCat.List, tReq)

	tReq.Url = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tReq.Surl = "ABCDEFGHIJ"
	tCat.List = append(tCat.List, tReq)

	tReq.Url = "abcdefghijklmnopqrstuvwxyz"
	tReq.Surl = "abcdefghij"
	tCat.List = append(tCat.List, tReq)

	tReq.Url = `!"#$%&'()*+,-./:;<=>?@"`
	tReq.Surl = "__________"
	tCat.List = append(tCat.List, tReq)

	tReq.Url = `[\]^_`
	tReq.Surl = "1111111112"
	tCat.List = append(tCat.List, tReq)

	tReq.Url = `{|}~¡¢£¤¥¦§¨©ª«¬­®¯°±²³´µ¶·¸¹º»¼½¾¿`
	tReq.Surl = "1111111113"
	tCat.List = append(tCat.List, tReq)

	tReq.Url = `ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïðñòóôõö÷øùúûüýþÿĀāĂăĄąĆćĈĉĊċČčĎďĐđĒēĔĕĖėĘęĚěĜĝĞğĠġĢģĤĥĦħĨĩĪīĬĭĮįİıĲĳĴĵĶķĸĹĺĻļĽľĿŀŁłŃńŅņŇňŉŊŋŌōŎŏŐőŒœŔŕŖŗŘřŚśŜŝŞşŠšŢţŤťŦŧŨũŪūŬŭŮůŰűŲųŴŵŶŷŸŹźŻżŽžſƀƁƂƃƄƅƆƇƈƉƊƋƌƍƎƏƐƑƒƓƔƕƖƗƘƙƚƛƜƝƞƟƠơƢƣƤƥƦƧƨƩƪƫƬƭƮƯưƱƲƳƴƵƶƷƸƹƺƻƼƽƾƿǀǁǂǃǄǅǆǇǈǉǊǋǌǍǎǏǐǑǒǓǔǕǖǗǘǙǚǛǜǝǞǟǠǡǢǣǤǥǦǧǨǩǪǫǬǭǮǯǰǱǲǳǴǵǶǷǸǹǺǻǼǽǾǿȀȁȂȃȄȅȆȇȈȉȊȋȌȍȎȏȐȑȒȓȔȕȖȗȘșȚțȜȝȞȟȠȡȢȣȤȥȦȧȨȩȪȫȬȭȮȯȰȱȲȳȴȵȶȷȸȹȺȻȼȽȾȿɀɁɂɃɄɅɆɇɈɉɊɋɌɍɎɏɐɑɒɓɔɕɖɗɘəɚɛɜɝɞɟɠɡɢɣɤɥɦɧɨɩɪɫɬɭɮɯɰɱɲɳɴɵɶɷɸɹɺɻɼɽɾɿʀʁʂʃʄʅʆʇʈʉʊʋʌʍʎʏʐʑʒʓʔʕʖʗʘʙʚʛʜʝʞʟʠʡʢʣʤʥʦʧʨʩʪʫʬʭʮ`
	tReq.Surl = "1111111114"
	tCat.List = append(tCat.List, tReq)
	return tCat
}
