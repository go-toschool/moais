package core 

type Rules struct {
	//rules or elements

	// component name
}

func (r *Rules) add () error {
	return nil
}

func (r *Rules) remove () error {
	return nil
}

//import * as arrayLocalFile from '../../helpers/array-local-files.methods'
// -> helpers thats operate with rules file

//define register mqtt topics
// add rule
// add massive
// remove rule
// remove all rules
// change rule
// get rules info

//timekeeper struct that dispatch event_name with scheduler and params info
/*
Timekeeper.on(TIMEKEEPER_DISPATCHER_EVENT_NAME, (scheduler: IAllSchedulersInfo) => {
      const rules = this.getByType(TYPES.SCHEDULER)

      rules.forEach((rule) => {
        if (rule.getActive()) {
          rule.getInfo()
            .then((info: IAllRulesInfo) => {
              const schedulers: string[] = info.params.schedulers

              if (lodash.includes(schedulers, scheduler.id)) {
                const ruleParams = Object.keys(info.params).reduce((obj, param) => {
                  if (param !== 'schedulers') {
                    obj[param] = info.params[param]
                  }
                  return obj
                }, {})

                const params = Object.assign({}, ruleParams, scheduler.params)
                rule.runFunction(params, {
                  wisebotId: wisebot.id,
                  talkative: wisebot.talkative,
                  devices: wisebot.devicesEngine,
                  globalValues: wisebot.globalValues,
                  logger: Logger,
                  hasInternet,
                  storage: StorageClient,
                }, this.callbackFunction)
              }
            })
            .catch((err) => Logger.error(`Could not get rule's info in RulesEngine#constructor (${TIMEKEEPER_DISPATCHER_EVENT_NAME})`, {rule_id: rule.getId()}))
        }
      })
    })
*/

//define event that use dispatcher by event_name (event params)
/*
Eventkeeper.on(EVENTKEEPER_DISPATCHER_EVENT_NAME, (event) => {
      const rules = this.getByType(TYPES.EVENT)

      rules.forEach((rule) => {
        if (rule.getActive()) {
          rule.getInfo()
            .then((info: IAllRulesInfo) => {
              const events: string[] = info.params.events

              events.forEach((ev) => {
                if (wildcard(ev, event.name)) {
                  rule.runFunction(event, {
                    wisebotId: wisebot.id,
                    talkative: wisebot.talkative,
                    devices: wisebot.devicesEngine,
                    globalValues: wisebot.globalValues,
                    logger: Logger,
                    hasInternet,
                  }, this.callbackFunction)
                }
              })
            })
            .catch((err) => Logger.error(`Could not get rule's info in RulesEngine#constructor (${EVENTKEEPER_DISPATCHER_EVENT_NAME})`, {rule_id: rule.getId()}))
        }
      })
    })
*/

//ENGINE METHODS
// public removeAll (): Promise<void> {
// public replace (element: T, index: number): void {
// public getAll (): T[] {
// public getById (id: string): T {
// public getByType (type: string): T[] {
// public getByName (name: string): T {

//instanceAndAddInLocal method
//add rule method
//add Massive rule method
//change rule method
//change massive rule method
//callback method that use on any rule lambda method
//getInfo that return all rules information
//halt methods that stop rule executation with repsectives precautions
