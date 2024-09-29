import GuestEntity from "@/booking/core/domain/entities/guest.entity";
import { IGuestRepository } from "@/booking/core/domain/ports/guest.repository";
import { PrismaService } from "../prisma.service";

export class GuestRepository implements IGuestRepository{
    constructor(private prisma: PrismaService){}
    
    Get(id: Number): GuestEntity {
        throw new Error("Method not implemented.");
    }
    
    async Create(guest: GuestEntity): Promise<number> {
        const data = guest.toJSON();

        const newGuest = await this.prisma.guest.create({
            data: {  
                idNumber: data.document.idNumber,
                documentType: data.document.documentType,
                ...data
            }
        });

        return newGuest.id;
    }
}