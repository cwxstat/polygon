package industry

import (
	"github.com/cwxstat/polygon/request"

	"strings"
)

var (
	url = "https://storage.googleapis.com/montco-stats/stocks/allstocks.csv"
)

type INDUSTRY struct {
	url      string
	values   [][]string
	sector   map[string][]string
	industry map[string][]string
}

func NewIndustry() *INDUSTRY {
	i := &INDUSTRY{url: url}
	i.sector = map[string][]string{}
	i.industry = map[string][]string{}
	return i
}

// FIXME: (mmc) better way
func splitLine(s string) []string {
	r := strings.Split(s, "\"")
	if len(r) == 3 {
		result := strings.Split(r[0], ",")
		result = result[0 : len(result)-1]
		result = append(result, r[1])
		t := strings.Split(r[2], ",")
		result = append(result, t[1:]...)
		return result

	}
	return strings.Split(r[0], ",")
}

func walk(s string) [][]string {
	result := [][]string{}
	for _, v := range strings.Split(s, "\r\n") {
		result = append(result, splitLine(v))
	}
	return result
}

func (i *INDUSTRY) rawList() ([][]string, error) {
	body, err := request.Request(i.url)
	return walk(body), err

}

func max(i []int) int {
	max := 0
	for _, v := range i {
		if v > max {
			max = v
		}
	}
	return max
}

/* collect:

     collect(list, [1,2,3,4,6], 1)
	 will skip first row 1, and only collect columns 1,2,3,4,6
*/
func collect(raw [][]string, indexes []int, skipRows int) [][]string {

	max := max(indexes)
	result := [][]string{}
	for k, v := range raw {
		// Heading or bad record
		if k < skipRows || len(v) < max {
			continue
		}
		t := []string{}
		for _, i := range indexes {
			t = append(t, v[i])
		}

		result = append(result, t)
	}
	return result
}

func (i *INDUSTRY) buildIndustry() ([][]string, error) {
	list, err := i.rawList()
	if err != nil {
		return list, nil
	}
	i.values = collect(list, []int{1, 2, 3, 4, 6}, 1)

	for _, v := range i.values {
		i.sector[v[2]] = append(i.sector[v[2]], v[0])
		i.industry[v[3]] = append(i.industry[v[3]], v[0])
	}

	return i.values, nil
}

func (i *INDUSTRY) Tickers(industry string) []string {

	if i.values == nil {
		i.buildIndustry()
	}

	return []string{}
}
