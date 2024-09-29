import GuestEntity from "@/booking/core/domain/entities/guest.entity";
import { GuestDTO } from "../dtos/guest.dto";

export class GuestRequest {
    constructor(public data: GuestDTO){}

    static mapToEntity(guestDTO: GuestDTO): GuestEntity{
        return new GuestEntity({
            email: guestDTO.email,
            name: guestDTO.name,
            surname: guestDTO.surname,
            id: guestDTO.id,
            document: {
                documentType: guestDTO.idTypeCode,
                idNumber: guestDTO.idNumber
            }
        });
    }
}