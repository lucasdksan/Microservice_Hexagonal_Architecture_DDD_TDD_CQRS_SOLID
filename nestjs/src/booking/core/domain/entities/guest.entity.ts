import { Entity } from "@/shared/domain/entities/entity";
import { Person } from "../value-objects/person.valueobject";

export type GuestProps = {
    id: number;
    name: string;
    surname: string;
    email: string;
    document: Person;
}

export default class GuestEntity extends Entity<GuestProps> {
    constructor(public readonly props: GuestProps){
        super(props);
    }
}