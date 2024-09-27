import GuestEntity from "../entities/guest.entity";

export interface IGuestRepository {
    Get(id: Number): GuestEntity;
    Save(guest: GuestEntity): number;
}