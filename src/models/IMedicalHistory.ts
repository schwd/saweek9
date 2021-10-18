import { MedicalRecordInterface } from "./IMedicalRecord";

export interface MedicalHistoryInterface {

    ID: string,
   
    DiseaseID: string;
   
    Diagnosis: string;
   
    Treatment: string;
   
    DepartmentID: string;

    DoctorID: string;
   
    Date: Date | null;

    MedicalRecordID: number,
    MedicalRecord:   MedicalRecordInterface,
   
   }