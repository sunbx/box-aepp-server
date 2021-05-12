package controllers

import (
	"box/models"
	"box/utils"
	"encoding/json"
	"strconv"
)

type OracleProblemController struct {
	BaseController
}
type OracleProblemInfoController struct {
	BaseController
}

func (c *OracleProblemInfoController) Post() {
	id := c.GetString("id")
	contractResult, _, e := models.CallStaticContractFunction("ak_idkx6m3bgRr7WiKXuB8EBYBoRqVsaSc6qo4dsd23HKgj3qiCF", models.OraclesContractV1, "get_problem", []string{id})
	if e != nil {
		c.SuccessJson(e.Error())
		return
	}
	switch contractResult.(type) {
	case map[string]interface{}:
		problem := contractResult.(map[string]interface{})
		answerResultArray := problem["answer"].([]interface{})

		var answers []interface{}
		for i := 0; i < len(answerResultArray); i++ {
			answerArray := answerResultArray[i].([]interface{})
			answer := answerArray[1].(map[string]interface{})
			answer["index"] = i
			answers = append(answers, answer)
		}
		problem["answer"] = answers
		count, _ := problem["count"].(json.Number).Float64()
		problem["count"] = utils.FormatTokens(count, -1)

		minCount, _ := problem["min_count"].(json.Number).Float64()
		problem["min_count"] = utils.FormatTokens(minCount, -1)
		index, _ := strconv.Atoi(id)
		problem["index"] = index + 1

		c.SuccessJson(problem)
		return
	}

	c.SuccessJson(contractResult)
}

func (c *OracleProblemController) Post() {
	t := c.GetString("type")
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
			problemArray := contractResultArray[i].([]interface{})
			problem := problemArray[1].(map[string]interface{})
			answerResultArray := problem["answer"].([]interface{})

			var answers []interface{}
			for i := 0; i < len(answerResultArray); i++ {
				answerArray := answerResultArray[i].([]interface{})
				answer := answerArray[1].(map[string]interface{})
				answer["index"] = i
				answers = append(answers, answer)
			}
			problem["answer"] = answers
			count, _ := problem["count"].(json.Number).Float64()
			problem["count"] = utils.FormatTokens(count, -1)

			minCount, _ := problem["min_count"].(json.Number).Float64()
			problem["min_count"] = utils.FormatTokens(minCount, -1)
			problem["index"] = i + 1

			status, _ := problem["status"].(json.Number).Float64()
			if t == "finish" && status != 0 {
				continue
			}

			problems = append(problems, problem)

		}
		if problems == nil {
			c.SuccessJson([]JsonData{})
			return
		}

		c.SuccessJson(problems)
		return
	}

	c.SuccessJson(contractResult)
}
