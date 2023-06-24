//Copyright © 2022 Ugo Landini <ugo.landini@gmail.com>
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in
//all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
//THE SOFTWARE.

package functions

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"github.com/ugol/jr/pkg/configuration"
	"github.com/ugol/jr/pkg/constants"
	"github.com/ugol/jr/pkg/ctx"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

func FunctionsMap() template.FuncMap {
	return fmap
}

var Random = rand.New(rand.NewSource(0))
var data = map[string][]string{}
var fmap = map[string]interface{}{

	// text utilities
	"atoi":                     strconv.Atoi,
	"counter":                  Counter,
	"first":                    func(s string) string { return s[:1] },
	"firstword":                func(s string) string { return strings.Split(s, " ")[0] },
	"from":                     Word,
	"from_at":                  WordAt,
	"from_shuffle":             WordShuffle,
	"from_n":                   WordShuffleN,
	"join":                     strings.Join,
	"len":                      Len,
	"lower":                    strings.ToLower,
	"lorem":                    Lorem,
	"markov":                   Nonsense,
	"random":                   func(s []string) string { return s[Random.Intn(len(s))] },
	"randoms":                  func(s string) string { a := strings.Split(s, "|"); return a[Random.Intn(len(a))] },
	"random_index":             RandomIndex,
	"random_string":            RandomString,
	"random_string_vocabulary": RandomStringVocabulary,
	"regex":                    Regex,
	"repeat":                   strings.Repeat,
	"replaceall":               strings.ReplaceAll,
	"sentence":                 Sentence,
	"sentence_prefix":          SentencePrefix,
	"squeeze":                  func(s string) string { return strings.ReplaceAll(s, " ", "") },
	"squeezechars":             func(s, c string) string { return strings.ReplaceAll(s, c, "") },
	"split":                    strings.Split,
	"substr":                   func(start, length int, s string) string { return s[start:length] },
	"trim":                     strings.TrimSpace,
	"trimchars":                strings.Trim,
	"title":                    cases.Title(language.English).String,
	"upper":                    strings.ToUpper,

	// math utilities
	"add":          func(a, b int) int { return a + b },
	"div":          func(a, b int) int { return a / b },
	"format_float": func(f string, v float32) string { return fmt.Sprintf(f, v) },
	"integer":      func(min, max int) int { return min + Random.Intn(max-min) },
	"integer64":    func(min, max int64) int64 { return min + Random.Int63n(max-min) },
	"floating":     func(min, max float32) float32 { return min + Random.Float32()*(max-min) },
	"sub":          func(a, b int) int { return a - b },
	"max":          math.Max,
	"min":          math.Min,
	"minint":       Minint,
	"maxint":       Maxint,
	"mod":          func(a, b int) int { return a % b },
	"mul":          func(a, b int) int { return a * b },

	// networking and time utilities
	"http_method":       HttpMethod,
	"ip":                Ip,
	"ipv6":              Ipv6,
	"ip_known_protocol": IpKnownProtocol,
	"ip_known_port":     IpKnownPort,
	"mac":               Mac,
	"password":          Password,
	"useragent":         UserAgent,

	// people related utilities
	"cf":             CodiceFiscale,
	"company":        Company,
	"email":          Email,
	"email_provider": EmailProvider,
	"email_work":     WorkEmail,
	"gender":         Gender,
	"middlename":     Middlename,
	"name":           Name,
	"name_m":         NameM,
	"name_f":         NameF,
	"ssn":            Ssn,
	"surname":        Surname,
	"user":           User,
	"username":       Username,

	// address
	"building":       BuildingNumber,
	"cardinal":       Cardinal,
	"capital":        Capital,
	"capital_at":     CapitalAt,
	"city":           City,
	"city_at":        CityAt,
	"country":        Country,
	"country_random": CountryRandom,
	"country_at":     CountryAt,
	"latitude":       Latitude,
	"longitude":      Longitude,
	"nearby_gps":     NearbyGPS,
	"state":          State,
	"state_at":       StateAt,
	"state_short":    StateShort,
	"state_short_at": StateShortAt,
	"street":         Street,
	"zip":            Zip,
	"zip_at":         ZipAt,

	//finance
	"account":      Account,
	"amount":       Amount,
	"bitcoin":      Bitcoin,
	"card":         CreditCard,
	"cardCVV":      CreditCardCVV,
	"cusip":        Cusip,
	"ethereum":     Ethereum,
	"isin":         Isin,
	"sedol":        Sedol,
	"stock_symbol": StockSymbol,
	"swift":        Swift,
	"valor":        Valor,
	"wkn":          Wkn,

	//time and dates
	"birthdate":       BirthDate,
	"date_between":    DateBetween,
	"dates_between":   DatesBetween,
	"future":          Future,
	"past":            Past,
	"recent":          Recent,
	"soon":            Soon,
	"unix_time_stamp": UnixTimeStamp,

	//phone
	"country_code":    CountryCode,
	"country_code_at": CountryCodeAt,
	"imei":            Imei,
	"phone":           Phone,
	"phone_at":        PhoneAt,
	"mobile_phone":    MobilePhone,
	"mobile_phone_at": MobilePhoneAt,

	// generic utilities
	"array":    func(count int) []int { return make([]int, count) },
	"bool":     RandomBool,
	"image":    Image,
	"image_of": ImageOf,
	"index_of": IndexOf,
	"key":      func(name string, n int) string { return fmt.Sprintf("%s%d", name, Random.Intn(n)) },
	"seed":     Seed,
	"uuid":     UniqueId,
	"yesorno":  YesOrNo,

	// context utilities
	"add_v_to_list":        AddValueToList,
	"random_v_from_list":   RandomValueFromList,
	"random_n_v_from_list": RandomNValuesFromList,
	"get_v":                GetV,
	"set_v":                SetV,
}

// Seed sets seeds and can be used in a template
func Seed(rndSeed int64) string {
	SetSeed(rndSeed)
	return ""
}

// SetSeed sets seeds for all random JR objects
func SetSeed(rndSeed int64) {
	Random.Seed(rndSeed)
	uuid.SetRand(Random)
}

// AddValueToList adds value v to Context list l
func AddValueToList(l string, v string) string {
	ctx.JrContext.CtxListLock.Lock()
	defer ctx.JrContext.CtxListLock.Unlock()
	ctx.JrContext.CtxList[l] = append(ctx.JrContext.CtxList[l], v)
	return ""
}

func ExtractMetaFrom(outTemplate string) (string, string) {
	start := strings.LastIndex(outTemplate, "_meta")
	if start == -1 {
		return "", outTemplate
	}
	end := strings.Index(outTemplate, "},")
	meta := outTemplate[start+7 : end+1]
	tpl := outTemplate[0:start-1] + outTemplate[end+2:]
	return meta, tpl
}

// GetV gets value s from Context
func GetV(s string) string {
	ctx.JrContext.CtxLock.RLock()
	defer ctx.JrContext.CtxLock.RUnlock()
	return ctx.JrContext.Ctx[s]
}

// SetV adds value v to Context
func SetV(s string, v string) string {
	ctx.JrContext.CtxLock.Lock()
	defer ctx.JrContext.CtxLock.Unlock()
	ctx.JrContext.Ctx[s] = v
	return ""
}

// IndexOf returns the index of the s string in a file
func IndexOf(s string, name string) int {
	_, err := Cache(name)
	if err != nil {
		return -1
	}
	words := data[name]
	index := sort.Search(len(words), func(i int) bool { return strings.ToLower(words[i]) >= strings.ToLower(s) })

	if index < len(words) && words[index] == s {
		return index
	} else {
		return -1
	}
}

// Len returns number of words (lines) in a word file
func Len(name string) string {
	_, err := Cache(name)
	if err != nil {
		return ""
	}
	l := len(data[name])
	return strconv.Itoa(l)
}

// RandomIndex returns a random index in a word file
func RandomIndex(name string) string {
	_, err := Cache(name)
	if err != nil {
		return ""
	}
	words := data[name]
	ctx.JrContext.LastIndex = Random.Intn(len(words))
	return strconv.Itoa(ctx.JrContext.LastIndex)
}

// RandomValueFromList returns a random value from Context list l
func RandomValueFromList(s string) string {
	ctx.JrContext.CtxListLock.RLock()
	defer ctx.JrContext.CtxListLock.RUnlock()
	list := ctx.JrContext.CtxList[s]
	l := len(list)
	if l != 0 {
		return list[Random.Intn(l)]
	} else {
		return ""
	}
}

// RandomNValuesFromList returns a random value from Context list l
func RandomNValuesFromList(s string, n int) []string {
	ctx.JrContext.CtxListLock.RLock()
	defer ctx.JrContext.CtxListLock.RUnlock()
	list := ctx.JrContext.CtxList[s]
	l := len(list)
	if l != 0 {
		ints := findNDifferentInts(n, l)
		results := make([]string, len(ints))
		for i := range ints {
			results[i] = list[i]
		}
		return results
	} else {
		return []string{""}
	}
}

// Word returns a random string from a list of strings in a file.
func Word(name string) string {
	_, err := Cache(name)
	if err != nil {
		return ""
	}
	words := data[name]
	ctx.JrContext.LastIndex = Random.Intn(len(words))
	return words[ctx.JrContext.LastIndex]
}

// WordAt returns a string at a given position in a list of strings in a file.
func WordAt(name string, index int) string {
	_, err := Cache(name)
	if err != nil {
		return ""
	}
	words := data[name]
	return words[index]
}

// WordShuffle returns a shuffled list of strings in a file.
func WordShuffle(name string) []string {
	_, err := Cache(name)
	if err != nil {
		return []string{""}
	}
	words := data[name]
	return WordShuffleN(name, len(words))
}

// wordShuffleN return a subset of n elements in a list of string in a file.
func WordShuffleN(name string, n int) []string {
	_, err := Cache(name)
	if err != nil {
		return []string{""}
	}
	words := data[name]
	Random.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	number := Minint(n, len(words))
	return words[:number]
}

// Minint returns the minimum between two ints
func Minint(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Maxint returns the minimum between two ints
func Maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Cache is used to internally Cache data from word files
func Cache(name string) (bool, error) {

	templateDir := configuration.GlobalCfg.TemplateDir
	if templateDir == "" {
		templateDir = fmt.Sprintf("%s/%s", constants.DEFAULT_HOMEDIR, "templates")
	}
	v := data[name]
	if v == nil {
		locale := strings.ToLower(ctx.JrContext.Locale)
		filename := fmt.Sprintf("%s/data/%s/%s", os.ExpandEnv(templateDir), locale, name)
		if locale != "us" && !(fileExists(filename)) {
			filename = fmt.Sprintf("%s/data/%s/%s", os.ExpandEnv(templateDir), "us", name)
		}
		data[name] = initialize(filename)
		if len(data[name]) == 0 {
			return false, fmt.Errorf("no words found in %s", filename)
		} else {
			return true, nil
		}
	}
	return false, nil
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else {
		return false
	}
}

// Helper function to generate n different integers from 0 to length
func findNDifferentInts(n, max int) []int {

	n = Minint(n, max)
	ints := make([]int, n)

	// Generate n different random indices of maximum length
	for i := 0; i < n; {
		index := Random.Intn(max)
		if !contains(ints, index) {
			ints[i] = index
			i++
		}
	}

	return ints
}

// Helper function to check if an int is in a slice of ints
func contains(values []int, value int) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func initialize(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Failed to open file: %s", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Error in closing file: %s", err)
		}
	}(file)

	var words []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words
}
