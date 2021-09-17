import { writable } from "svelte/store"
import { buildValidator } from "./validate"

export function createFieldValidator(...validators) {
    const { subscribe, set } = writable({ dirty: false, valid: false, message: null })
    const validator = buildValidator(validators)

    function action(node, binding) {
        function validate(value, dirty) {
            const result = validator(value, dirty)
            set(result)
        }

        validate(binding, false)

        return {
            update(value) {
                validate(value, true)
            }
        }
    }

    function reset() {
        setTimeout(() => {
            set({ dirty: false, valid: false, message: null })
        }, 500);
    }

    return [{ subscribe }, action, reset]
}