import { AppResponse } from "@/shared/application/responses/response";
import { IGuestRepository } from "../domain/ports/guest.repository";
import { IGuestManager } from "./guest/ports/iguest-manager.port";
import { GuestRequest } from "./guest/requests/guest.request";
import { GuestResponse } from "./guest/responses/guest.response";

export class GuestManager implements IGuestManager {
    constructor(private readonly guestRepository: IGuestRepository){}

    async Create(request: GuestRequest): Promise<GuestResponse> {
        try {
            const guest = GuestRequest.mapToEntity(request.data);
            request.data.id = await this.guestRepository.Create(guest);

            return {
                data: request.data,
                success: true,
                message: ""
            };
        } catch(err) {
            return {
                success: false,
                message: "There was an error when saving to DB",
                errorCode: AppResponse.ErrorCodes.COULDNOT_STORE_DATA
            }
        }
    }
}