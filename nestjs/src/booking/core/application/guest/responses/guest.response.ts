import { AppResponse } from "@/shared/application/responses/response";
import { GuestDTO } from "../dtos/guest.dto";

export class GuestResponse extends AppResponse.Response {
    data?: GuestDTO;
}