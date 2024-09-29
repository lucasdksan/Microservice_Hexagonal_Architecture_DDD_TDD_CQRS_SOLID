import { Entity } from "@/shared/domain/entities/entity";
import { Price } from "../value-objects/price.valueobject";

export type RoomProps = {
    id: number;
    name: string;
    level: number;
    inMaintenance: boolean;
    price: Price;
}

export default class RoomEntity extends Entity<RoomProps> {
    constructor(public readonly props: RoomProps) {
        super(props);
    }

    isAvailable(): boolean {
        return !this.props.inMaintenance || this.hasGuest() ? false : true;
    }

    hasGuest(): boolean {
        return true;
    }
}