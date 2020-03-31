package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Scavenge is the Scavenge struct
type Articles struct {
	Article_id   string          `json:"article_id"`                     // id of the article
	Uid          string          `json:"uid"`							 // id of the user
	Tid          string          `json:"tid"`                            // id of the transaction
	A_timestamp  string          `json:"a_timestamp"`                    // timestamp of the article 
	A_title      string          `json:"a_title"`                        // title of the article 
	A_text       string          `json:"a_text"`                          // text of the article
	TextHash     string          `json:"textHash"`
	Tag          string          `json:"tag"` 
	Flag         int             `json:"flag"` 
}  


type Articlevote struct {
	Article_id   string          `json:"article_id"`                     // id of the article
	VoteUP       int             `json:"voteUP"`
	VoteDOWN     int             `json:"voteDOWN"`
	Num          int             `json:"num"`
} 


type Comment struct {
	Comment_id   string          `json:"comment_id"`                     // id of the comment
	Article_id   string          `json:"article_id"`                     // id of the article
	Tid          string          `json:"tid"`                            // id of the transaction
	Uid          string          `json:"uid"`							 // id of the user
	C_timestamp  string          `json:"c_timestamp"`                    // timestamp of the comment 
	C_text       string          `json:"c_text"`                         // context of the comment
	Flag         int             `json:"flag"`
} 


type Commentvote struct {
	Comment_id   string          `json:"comment_id"`                     // id of the comment
	VoteUP       int             `json:"voteUP"`
	VoteDOWN     int             `json:"voteDOWN"`
	Num          int             `json:"num"`
} 


type Domain struct {
	Domainn      string             `json:"domain"`                         // 
	Ip           string             `json:"ip"`
	Owner        string             `json:"owner"`
	Suffix       string             `json:"suffix"`
} 


type Equity struct {
	Uid          string             `json:"uid"`
	Balance      string             `json:"balance"`                        // 股权资产数
	Detail       string             `json:"detail"`
}


type EquityTransaction struct {
	Et_id          string            `json:"et_id"`
	Source_id      string            `json:"source_id"`
	Destination_id string            `json:"destination_id"`
	Tid            string            `json:"tid"`
	Balance        string            `json:"balance"`
	Et_timestamp   int               `json:"et_timestamp"`
	Detail         string            `json:"detail"`
}

type ReturnVisit struct{
	Return_visit_id  string           `json:"return_visit_id"`
	Article_id       string           `json:"article_id"`
	Tid              string           `json:"tid"`
	Uid              string           `json:"uid"`
	Rv_timestamp     string           `json:"rv_timestamp"`
	Rv_text			 string           `json:"rv_text"`
	Flag             string           `json:"flag"` 
}


// implement fmt.Stringer
func (s Scavenge) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Creator: %s
	Description: %s
	SolutionHash: %s
	Reward: %s
	Solution: %s
	Scavenger: %s`,
		s.Creator,
		s.Description,
		s.SolutionHash,
		s.Reward,
		s.Solution,
		s.Scavenger,
	))
}

// Commit is the commit struct
type Commit struct {
	Scavenger             sdk.AccAddress `json:"scavenger" yaml:"scavenger"`                         // address of the scavenger scavenger
	SolutionHash          string         `json:"solutionHash" yaml:"solutionHash"`                   // SolutionHash of the scavenge
	SolutionScavengerHash string         `json:"solutionScavengerHash" yaml:"solutionScavengerHash"` // solution hash of the scavenge
}

// implement fmt.Stringer
func (c Commit) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Scavenger: %s
	SolutionHash: %s
	SolutionScavengerHash: %s`,
		c.Scavenger,
		c.SolutionHash,
		c.SolutionScavengerHash,
	))
}
