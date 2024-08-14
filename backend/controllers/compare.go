package controllers

import (
	"container/list"
	"encoding/json"
	"math"
	"math/rand"
	"net/http"
)

type Weapon struct {
  Range int `json:"range"`
  Attacks int `json:"attacks"`
  ToHit int `json:"toHit"`
  ToWound int `json:"toWound"`
  Rend int `json:"rend"`
  Damage int `json:"damage"`
  Models int `json:"models"`
}

type Unit struct {
  Name string `json:"name"`
  Health int `json:"health"`
  Save int `json:"save"`
  Weapons []Weapon `json:"weapons"`
}

var TestUnit = Unit{
  Name: "Test Unit",
  Health: 10,
  Save: 4,
  Weapons: []Weapon{
    {
      Range: 1,
      Attacks: 1,
      ToHit: 3,
      ToWound: 3,
      Rend: 0,
      Damage: 1,
      Models: 10,
    },
  },
}

var TestUnit2 = Unit{
  Name: "Test Unit 2",
  Health: 10,
  Save: 4,
  Weapons: []Weapon{
    {
      Range: 1,
      Attacks: 2,
      ToHit: 4,
      ToWound: 4,
      Rend: 1,
      Damage: 1,
      Models: 10,
    },
  },
}

type Dice struct {
  Sides int `json:"sides"`
}


func (d *Dice) Roll() float64 {
  return math.Floor(rand.Float64() * float64(d.Sides)) + 1
}

func (*Dice) Average() float64 {
  return 3.5
}

func (u *Unit) Attack(target *Unit)map[string]interface{} {
  dice := Dice{Sides: 6}
  weaponList := list.New()
  originalHealth := target.Health
  for _, weapon := range u.Weapons {
    weaponList.PushBack(weapon)
  }
  for e := weaponList.Front(); e != nil; e = e.Next() {
    weapon := e.Value.(Weapon)
    for i := 0; i < weapon.Attacks; i++ {
      hitRoll := int(dice.Roll())
      woundRoll := int(dice.Roll())
      if (hitRoll < weapon.ToHit) {
        continue
      }
      if (woundRoll < weapon.ToWound) {
        continue
      }
      damage := (target.Save - weapon.Rend) - weapon.Damage
      target.Health -= damage
    }
      weaponList.Remove(e)
    }
    return map[string]interface{}{
      "remainingHealth": target.Health,
      "damageDealt": originalHealth - target.Health,
    }
}



func CompareUnits() http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    attacker := TestUnit
    defender := TestUnit2
    simulationCount := 100000
    offense := attacker.Attack(&defender)
    defense := defender.Attack(&attacker)
    averageBattle := CompareAverage(attacker, defender, simulationCount)
    w.Header().Set("Content-Type", "application/json")
    data := map[string]interface{}{
      "attacker": offense,
      "defender": defense,
      "average": averageBattle,
    }
    json.NewEncoder(w).Encode(data)
  })
}

func CompareAverage(attacker Unit, defender Unit, simulations int)map[string]interface{} {
  attackerDamage := 0
  defenderDamage := 0
  attackRemainingHealth := 0
  defenseRemainingHealth := 0
  for i := 0; i < simulations; i++ {
    attackerCopy := attacker
    defenderCopy := defender
    offense := attackerCopy.Attack(&defenderCopy)
    defense := defenderCopy.Attack(&attackerCopy)
    attackerDamage += offense["damageDealt"].(int)
    defenderDamage += defense["damageDealt"].(int)
    attackRemainingHealth += defense["remainingHealth"].(int)
    defenseRemainingHealth += offense["remainingHealth"].(int)
  }
  return map[string]interface{}{
    "attacker": map[string]interface{}{
      "averageDamage": math.Ceil(float64(attackerDamage) / float64(simulations)),
      "averageRemainingHealth": math.Ceil(float64(attackRemainingHealth) / float64(simulations)),
    },
    "defender": map[string]interface{}{
      "averageDamage": math.Ceil(float64(defenderDamage) / float64(simulations)),
      "averageRemainingHealth": math.Ceil(float64(defenseRemainingHealth) / float64(simulations)),
    },
  }
}












