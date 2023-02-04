package repository

import "goTestFiber/model"

func (t *testRepo) Sum() model.Respond {
	var result model.Respond
	t.db.Table("test").Select("sum(active_power) as active_power, sum(power_input) as power_input").Scan(&result)
	return result
}
