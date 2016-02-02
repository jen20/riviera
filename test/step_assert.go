package test

import (
	"fmt"
	"log"
	"reflect"
)

type AssertFunc func(AzureStateBag) error

type StepAssert struct {
	StateBagKey string
	Checks      []AssertFunc
}

func (s *StepAssert) Run(state AzureStateBag) StepAction {
	hasErrors := false

	for _, v := range s.Checks {
		if err := v(state); err != nil {
			state.AppendError(err)
			hasErrors = true
		}
	}

	if hasErrors {
		return Halt
	}

	return Continue
}

func (s *StepAssert) Cleanup(state AzureStateBag) {
}

func CheckStringProperty(stateBagKey, propertyName, expectedValue string) AssertFunc {
	return func(state AzureStateBag) error {

		log.Printf("[INFO] Asserting %s.%s has value %q", stateBagKey, propertyName, expectedValue)

		stateValue, ok := state.GetOk(stateBagKey)
		if !ok {
			return fmt.Errorf("Internal Test Error - Cannot find state key %q in state", stateBagKey)
		}

		var v reflect.Value
		if reflect.ValueOf(stateValue).Kind() == reflect.Ptr {
			v = reflect.ValueOf(stateValue).Elem()
		} else {
			v = reflect.ValueOf(stateValue)
		}

		switch v.Kind() {
		case reflect.Struct:
			propertyField := v.FieldByName(propertyName)

			switch propertyField.Kind() {
			case reflect.String:
				actualValue := propertyField.String()
				if actualValue != expectedValue {
					return fmt.Errorf("%s.%s: Expected %q, Got %q", stateBagKey, propertyName, expectedValue, actualValue)
				}

				return nil

			case reflect.Ptr:
				actualValue := propertyField.Elem().String()
				if actualValue != expectedValue {
					return fmt.Errorf("%s: Expected %q, Got %q", propertyName, expectedValue, actualValue)
				}

				return nil

			default:
				return fmt.Errorf("Internal Test Error - %q is not a string or *string - checkStringProperty may not be used with this type %q", propertyName, propertyField.Kind())
			}

		default:
			return fmt.Errorf("Internal Test Error - Value for state bag key %q is not a struct", stateBagKey)
		}
	}
}
