import { RoomInterface } from './IRoom';
import { FoodOrderedsInterface } from './IFoodordered';
import { UserInterface } from './IUser';
export interface BookingsInterface {
    ID: number;
    BookingTimeStart: Date;
    BookingTimeStop: Date;
    RoomID: number;
    Room: RoomInterface;
    TotalPrice: number;
    
    MemberID: number;
    Member: UserInterface;

    FoodOrdereds: FoodOrderedsInterface[];
    
    
}