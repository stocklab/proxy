package detail

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"proxy"
	"strconv"
	"sync"
)

type data5u struct {
	*sync.RWMutex
	urlList     []string
	addressList []*proxy.Address
}

func init() {
	d := newData()
	d.Register(proxy.GlobalProxyManger)
}

func newData() *data5u {
	d := &data5u{urlList: []string {"http://www.data5u.com/"}}
	d.addressList = make([]*proxy.Address, 0, 2)
	return d
}

func (p *data5u) makeAddress(url, port, level, country string) (*proxy.Address, error) {
	addr := &proxy.Address{
		IP:     url,
		Country: country,
	}
	var err error
	addr.Port, err = strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	switch level {
	case "透明":
		addr.Level = proxy.Transparent
	case "匿名":
		addr.Level = proxy.Anonymous
	case "高匿":
		addr.Level = proxy.HighAnonymous
	default:
		addr.Level = proxy.Unknown
	}
	return addr, nil
}

func (p *data5u) run(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", proxy.GlobalUserAgent.Next())
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8")
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}
	addressList := make([]*proxy.Address, 0)
	doc.Find("ul[class=l2]").Each(func (i int, s *goquery.Selection) {
		data := make([]string, 0, 4)
		s.Find("li").Each(func (j int, li *goquery.Selection) {
			if j < 4 {
				data = append(data, li.Text())
			}
		})
		if len(data) == 4 {
			addr, err := p.makeAddress(data[0], data[1], data[2], data[3])
			if err == nil {
				addressList = append(addressList, addr)
			}
		}
	})
	if len(addressList) > 0 {
		p.Lock()
		defer p.Unlock()
		p.addressList = append(p.addressList, addressList...)
	}
	return nil
}

func (p *data5u) Run() error {
	for _, url := range p.urlList {
		p.run(url)
	}
	return nil
}

func (p *data5u) AddressList() []*proxy.Address {
	p.Lock()
	defer p.Unlock()
	address := p.addressList
	p.addressList = make([]*proxy.Address, 0, 2)
	return address
}

func (p *data5u) Register(m *proxy.Manager) {
	m.Register(p)
}
