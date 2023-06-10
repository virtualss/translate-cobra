package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const supportTranslate = `
 Name		Desc			Support Languages
"finance"	"金融财经领域"		"zh->en en->zh"
"machinery"	"机械制造领域"		"zh->en en->zh"
"senimed"	"生物医药领域"		"zh->en en->zh"
"novel"		"网络文学领域"		"zh->en"
"academic"	"学术论文领域"		"zh->en en->zh"
"novel"		"网络文学领域"		"zh->en en->zh"
"aerospace"	"航空航天领域"		"zh->en en->zh"
"wiki"		"人文社科领域"		"zh->en       "
"news"		"新闻资讯领域"		"zh->en en->zh"
"contract"	"合同领域"		"zh->en en->zh"
`

func NewField() *cobra.Command {
	filedCmd := &cobra.Command{
		Use:   "field",
		Short: "support domains",
		Long:  "support domains that could be use in 'from' and 'to'",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(supportTranslate)
		},
		//Example: "tl field",
	}
	return filedCmd
}
