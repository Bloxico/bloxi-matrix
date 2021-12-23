import { Test } from './test'
import { PercentageSet } from './shared'
import { Feature } from './feature'
import { SystemScore } from './systemScore'

// System is a top-level (architectural) denomination of a software project.
export class System {
  constructor(
    public name: string,
    public testKindStats: PercentageSet,
    public testStatusStats: PercentageSet,
    public score: SystemScore,
    public subsystems: SubSystem[] = [],
  ) {}

  public get id(): string {
    return this.name
  }
}

// Subsystem is a child of a System - second level (architectural) denomination of a software project
export class SubSystem extends System {
  constructor(
    public system: string,
    public features: Feature[],
    public tests: Test[],

    ...p: ConstructorParameters<typeof System>
  ) {
    super(...p)
  }

  public get id(): string {
    return `${this.system}/${this.name}`
  }
}
