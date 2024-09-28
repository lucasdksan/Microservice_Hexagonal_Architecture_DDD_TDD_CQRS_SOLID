import BookingEntity from "@/booking/core/domain/entities/booking.entity";
import { Actions } from "@/booking/core/domain/enums/actions.enum";
import { Status } from "@/booking/core/domain/enums/status.enum";

describe("Tests on the booking entity", ()=>{
    let booking:BookingEntity;

    beforeEach(()=> {
        booking = new BookingEntity();
    });

    it("Should Always Start With Created Status", ()=>{
        expect(booking.currentStatus()).toEqual(Status.Created);
    });

    it("Should set status to paid when paying for a booking with created status", ()=>{
        booking.changeState(Actions.Pay);
        
        expect(booking.currentStatus()).toEqual(Status.Paid);
    });

    it("Should set status to Canceld When Canceling a booking with Created Status", ()=>{
        booking.changeState(Actions.Cancel);
        
        expect(booking.currentStatus()).toEqual(Status.Canceled);
    });

    it("Should set status to finished When finishing a booking with Created Status", ()=>{
        booking.changeState(Actions.Pay);
        booking.changeState(Actions.Finish);
        
        expect(booking.currentStatus()).toEqual(Status.Finished);
    });

    it("Should set staus to refounded when finishing a paiid booking", ()=>{
        booking.changeState(Actions.Pay);
        booking.changeState(Actions.Refound);
        
        expect(booking.currentStatus()).toEqual(Status.Refounded);
    });

    it("Should set status to created when reopening a canceled booking", ()=>{
        booking.changeState(Actions.Cancel);
        booking.changeState(Actions.Reopen);
        
        expect(booking.currentStatus()).toEqual(Status.Created);
    });

    it("Shouldn't change status when refounding a booking with created status", ()=>{
        booking.changeState(Actions.Refound);
        
        expect(booking.currentStatus()).toEqual(Status.Created);
    });
    
    it("Shouldn't change status when refounding a finished booking", ()=>{
        booking.changeState(Actions.Pay);
        booking.changeState(Actions.Finish);
        booking.changeState(Actions.Refound);
        
        expect(booking.currentStatus()).toEqual(Status.Finished);
    });
});