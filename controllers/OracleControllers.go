package controllers

import (
	"box/models"
	"box/utils"
	"encoding/json"
)

type OracleProblemController struct {
	BaseController
}


func (c *OracleProblemController) Post() {
	//ctId := c.GetString("ct_id")
	contractResult, _, e := models.CallStaticContractFunction("ak_idkx6m3bgRr7WiKXuB8EBYBoRqVsaSc6qo4dsd23HKgj3qiCF", models.OraclesContractV1, "get_problems", []string{})
	if e != nil {
		c.SuccessJson(e.Error())
		return
	}
	switch contractResult.(type) {
	case []interface{}:
		contractResultArray := contractResult.([]interface{})
		var problems []interface{}
		for i := 0; i < len(contractResultArray); i++ {
			problemArray:= contractResultArray[i].([]interface{})
			problem := problemArray[1].(map[string]interface{})
			answerResultArray := problem["answer"].([]interface{})

			var answers []interface{}
			for i := 0; i < len(answerResultArray); i++ {
				answerArray:= answerResultArray[i].([]interface{})
				answer := answerArray[1].(map[string]interface{})
				answers = append(answers, answer)
			}
			problem["answer"] = answers
			count ,_:= problem["count"].(json.Number).Float64()
			problem["count"] = utils.FormatTokens(count, -1)

			minCount,_:= problem["min_count"].(json.Number).Float64()
			problem["min_count"] = utils.FormatTokens(minCount, -1)


			problems = append(problems, problem)



		}
		c.SuccessJson(problems)
		return
	}

	c.SuccessJson(contractResult)
}