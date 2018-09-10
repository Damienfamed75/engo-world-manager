package worldmanager

import (
	"fmt"
	"reflect"

	"engo.io/ecs"
)

// RemoveSystem will run through the world systems and remove
// the desired system you passed along with your parameters
func RemoveSystem(world *ecs.World, sysType interface{}, param ecs.BasicEntity) {
	for _, system := range world.Systems() {
		switch sys := reflect.TypeOf(system); sys {
		case reflect.TypeOf(sysType):
			system.Remove(param)
			fmt.Println("REMOVED")
		default:
			fmt.Println("FAILED")
		}
	}
}
