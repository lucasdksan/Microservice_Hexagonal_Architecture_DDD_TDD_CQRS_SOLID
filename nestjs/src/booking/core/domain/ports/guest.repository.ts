import GuestEntity from "../entities/guest.entity";

export interface IGuestRepository {
    Get(id: Number): GuestEntity;
    Create(guest: GuestEntity): Promise<number>;
}