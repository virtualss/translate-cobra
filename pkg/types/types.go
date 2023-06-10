package types

import (
	"strings"
)

type (
	TranslateConfig struct {
		AppID  string `json:"appid"`
		Secret string `json:"secret"`
		Domain string `json:"domain"`
		From   string `json:"from"`
		To     string `json:"to"`
	}
	Language struct {
		from string
		to   string
	}

	Domain struct {
		Name      string
		Desc      string
		Languages []*Language
	}

	DomainList   []*Domain
	LanguageList []*Language
)

func (d *DomainList) Support(domain string) (*Domain, bool) {
	for _, v := range *d {
		return v, strings.EqualFold(v.Name, domain)
	}
	return nil, false
}

func (l *LanguageList) Support(domain *Domain, from, to string) bool {
	for _, v := range domain.Languages {
		return *v == Language{from: from, to: to}
	}
	return false
}

var (
	ze = &Language{from: "zh", to: "en"}
	ez = &Language{from: "en", to: "zh"}

	it = &Domain{Name: "it", Desc: "信息技术领域", Languages: []*Language{ze, ez}}

	finance   = &Domain{Name: "finance", Desc: "金融财经领域", Languages: []*Language{ze, ez}}
	machinery = &Domain{Name: "machinery", Desc: "机械制造领域", Languages: []*Language{ze, ez}}
	senimed   = &Domain{Name: "senimed", Desc: "生物医药领域", Languages: []*Language{ze, ez}}
	novel     = &Domain{Name: "novel", Desc: "网络文学领域", Languages: []*Language{ze}}
	academic  = &Domain{Name: "academic", Desc: "学术论文领域", Languages: []*Language{ze, ez}}
	aerospace = &Domain{Name: "aerospace", Desc: "航空航天领域", Languages: []*Language{ze, ez}}
	wiki      = &Domain{Name: "wiki", Desc: "人文社科领域", Languages: []*Language{ze}}
	news      = &Domain{Name: "news", Desc: "新闻资讯领域", Languages: []*Language{ze, ez}}
	law       = &Domain{Name: "law", Desc: "法律法规领域", Languages: []*Language{ze, ez}}
	contract  = &Domain{Name: "contract", Desc: "合同领域", Languages: []*Language{ze, ez}}
)

var (
	Domains = DomainList{
		it,
		finance,
		machinery,
		senimed,
		novel,
		academic,
		aerospace,
		wiki,
		news,
		law,
		contract,
	}
	Languages = LanguageList{
		ze,
		ez,
	}
)

/*var domains []*DomainList = &DomainList{
	&Domain{Name: "it", Desc: "信息技术领域", Language: "zh->en en->zh"},
	&Domain{Name: "finance", Desc: "金融财经领域", Language: "zh->en en->zh"},
	&Domain{Name: "machinery", Desc: "机械制造领域", Language: "zh->en en->zh"},
	&Domain{Name: "senimed", Desc: "生物医药领域", Language: "zh->en en->zh"},
	&Domain{Name: "novel", Desc: "网络文学领域", Language: "zh->en"},
	&Domain{Name: "academic", Desc: "学术论文领域", Language: "zh->en en->zh"},
	&Domain{Name: "aerospace", Desc: "航空航天领域", Language: "zh->en en->zh"},
	&Domain{Name: "wiki", Desc: "人文社科领域", Language: "zh->en"},
	&Domain{Name: "news", Desc: "新闻资讯领域", Language: "zh->en en->zh"},
	&Domain{Name: "law", Desc: "法律法规领域", Language: "zh->en en->zh"},
	&Domain{Name: "contract", Desc: "合同领域", Language: "zh->en en->zh"}
}*/
